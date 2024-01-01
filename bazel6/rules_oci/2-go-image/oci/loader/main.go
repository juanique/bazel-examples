package main

import (
	"context"
	"fmt"
	"os"
)

func buildAndLoadImage(i Image, repoTags []string) error {
	dockerImageId := i.manifest.Config.Digest

	builder := NewImageBuilder(dockerImageId, repoTags)
	tarPath, err := builder.Build(i)
	if err != nil {
		return err
	}

	loader, err := NewDockerLoader()
	if err != nil {
		return err
	}

	action, err := loader.LoadTarIntoDocker(context.Background(), tarPath, dockerImageId, repoTags)

	if action.AlreadyLoaded {
		fmt.Println("Image ID", dockerImageId, "was already loaded.")
	}

	for _, tag := range action.TagsAlreadyPresent {
		fmt.Println("Image was already tagged with", tag)
	}

	for _, tag := range action.TagsAdded {
		fmt.Println("Tagged image with", tag)
	}

	if err != nil {
		return err
	}

	return nil
}

func main() {
	imagePath := os.Args[1]
	repoTags := os.Args[2:]

	image, err := NewImage(imagePath)
	if err != nil {
		panic(err)
	}
	err = buildAndLoadImage(image, repoTags)
	if err != nil {
		panic(err)
	}
}
