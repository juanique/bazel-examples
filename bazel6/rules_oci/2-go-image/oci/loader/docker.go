package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerLoadAction struct {
	AlreadyLoaded      bool
	TagsAdded          []string
	TagsAlreadyPresent []string
}

// DockerLoader holds a Docker client and provides methods to interact with Docker
type DockerLoader struct {
	cli *client.Client
}

// NewDockerLoader creates a new DockerLoader
func NewDockerLoader() (*DockerLoader, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("error creating Docker client: %w", err)
	}
	return &DockerLoader{cli: cli}, nil
}

// TagImage tags a Docker image with a new tag
func (d *DockerLoader) TagImage(ctx context.Context, imageID, tag string) error {
	err := d.cli.ImageTag(ctx, imageID, tag)
	if err != nil {
		return fmt.Errorf("error tagging image: %w", err)
	}
	return nil
}

// checkForExistingImage checks if an image with the specified ID exists in Docker
func (d *DockerLoader) checkForExistingImage(ctx context.Context, imageID string, tags []string) (DockerLoadAction, error) {
	action := DockerLoadAction{}

	images, err := d.cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return action, fmt.Errorf("error listing Docker images: %w", err)
	}

	tagsPresent := map[string]bool{}
	for _, tag := range tags {
		tagsPresent[tag] = false
	}

	var existingImage types.ImageSummary
	for _, image := range images {
		if image.ID == imageID {
			existingImage = image
			action.AlreadyLoaded = true
			break
		}
	}

	if !action.AlreadyLoaded {
		// We'll add all tags during the load itself
		action.TagsAdded = tags
		return action, nil
	}

	// The image was already there, we need to check if any extra tags are needed
	for _, tag := range existingImage.RepoTags {
		_, expected := tagsPresent[tag]
		if expected {
			tagsPresent[tag] = true
		}
	}

	for tag, alreadyPresent := range tagsPresent {
		if alreadyPresent {
			action.TagsAlreadyPresent = append(action.TagsAlreadyPresent, tag)
			continue
		}

		// Tag not there, we need to tag the image
		d.TagImage(ctx, imageID, tag)
		action.TagsAdded = append(action.TagsAlreadyPresent, tag)
	}

	return action, nil
}

// LoadTarIntoDocker loads a tar file into Docker if an image with the same ID doesn't exist
func (d *DockerLoader) LoadTarIntoDocker(ctx context.Context, tarPath, imageID string, repoTags []string) (DockerLoadAction, error) {
	// Check if the image already exists
	action, err := d.checkForExistingImage(ctx, imageID, repoTags)
	if err != nil {
		return action, err
	}
	if action.AlreadyLoaded {
		return action, nil
	}

	// Open the tar file
	tar, err := os.Open(tarPath)
	if err != nil {
		return action, fmt.Errorf("error opening tar file (%s): %w", tarPath, err)
	}
	defer tar.Close()

	// Load the tar file into Docker
	response, err := d.cli.ImageLoad(ctx, tar, true)
	if err != nil {
		return action, fmt.Errorf("error loading tar file into Docker: %w", err)
	}
	defer response.Body.Close()

	return action, nil
}
