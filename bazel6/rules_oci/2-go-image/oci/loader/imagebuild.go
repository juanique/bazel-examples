package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var acceptedMediaTypes = map[string]bool{
	"application/vnd.oci.image.manifest.v1+json":           true,
	"application/vnd.docker.distribution.manifest.v2+json": true,
}

type Descriptor struct {
	MediaType string `json:"mediaType"`
	Size      int    `json:"size"`
	Digest    string `json:"digest"`
}

type Manifest struct {
	MediaType string       `json:"mediaType"`
	Size      int          `json:"size"`
	Digest    string       `json:"digest"`
	Config    Descriptor   `json:"config"`
	Layers    []Descriptor `json:"layers"`
}

type ImageIndex struct {
	SchemaVersion int        `json:"schemaVersion"`
	MediaType     string     `json:"mediaType"`
	Manifests     []Manifest `json:"manifests"`
}

type Image struct {
	Path string

	// Dynamically loaded
	Index    ImageIndex
	manifest Manifest
}

func (i Image) BlobPath(digest string) string {
	return filepath.Join(i.Path, "blobs", strings.Replace(digest, ":", "/", -1))
}

func (i Image) IndexPath() string {
	return filepath.Join(i.Path, "index.json")
}

func (i Image) ConfigBlobPath() string {
	return i.BlobPath(i.manifest.Config.Digest)
}

func (i Image) ManifestBlobPath() string {
	return i.BlobPath(i.Index.Manifests[0].Digest)
}

func (i Image) GetLayerBlobPaths() []string {
	output := []string{}
	for _, layer := range i.manifest.Layers {
		output = append(output, i.BlobPath(layer.Digest))
	}

	return output
}

func (i *Image) LoadManifest() error {
	return loadJSON(i.ManifestBlobPath(), &i.manifest)
}

func (i *Image) LoadIndex() error {
	return loadJSON(i.IndexPath(), &i.Index)
}

func NewImage(path string) (Image, error) {
	image := Image{Path: path}
	if err := image.LoadIndex(); err != nil {
		return Image{}, err
	}
	if err := image.LoadManifest(); err != nil {
		return Image{}, err
	}

	return image, nil
}

type OutputFile struct {
	src string
	dst string
	rel string
}

type OutputManifest struct {
	Config   string   `json:"Config"`
	RepoTags []string `json:"RepoTags"`
	Layers   []string `json:"Layers"`
}

type ImageBuilder struct {
	stagingDir     string
	blobsDir       string
	outputManifest OutputManifest
	repoTags       []string

	// Stateful
	filesToCopy []OutputFile
}

func (b *ImageBuilder) Build(i Image) (string, error) {
	err := os.MkdirAll(b.blobsDir, 0o755)
	if err != nil {
		return "", fmt.Errorf("Failed to create output dir: %w", err)
	}

	if !acceptedMediaTypes[i.Index.Manifests[0].MediaType] {
		return "", fmt.Errorf("Unsupported media type: %s", i.Index.Manifests[0].MediaType)
	}

	b.outputManifest.RepoTags = b.repoTags

	configOutput := b.AddBlob(i.ConfigBlobPath())
	b.outputManifest.Config = configOutput.rel
	for _, layer := range i.GetLayerBlobPaths() {
		output := b.AddGZippedBlob(layer)
		b.outputManifest.Layers = append(b.outputManifest.Layers, output.rel)
	}

	tarInputs := []string{}
	for _, file := range b.filesToCopy {
		if err := copyFile(file.src, file.dst); err != nil {
			return "", err
		}
		tarInputs = append(tarInputs, file.rel)
	}

	outputManifest := []OutputManifest{b.outputManifest}
	err = WriteJSONToFile(outputManifest, b.GetOutputPath("manifest.json"))
	tarInputs = append(tarInputs, "manifest.json")
	if err != nil {
		return "", err
	}

	if err := createTar(b.stagingDir, b.GetOutputPath("image.tar"), tarInputs); err != nil {
		return "", fmt.Errorf("failed to create tar: %w", err)
	}

	return b.GetOutputPath("image.tar"), nil
}

func (b ImageBuilder) GetOutputPath(relPath string) string {
	return filepath.Join(b.stagingDir, relPath)
}

func (b *ImageBuilder) AddBlob(blobPath string) OutputFile {
	f := OutputFile{
		src: blobPath,
		dst: filepath.Join(b.blobsDir, filepath.Base(blobPath)),
	}
	f.rel, _ = filepath.Rel(b.stagingDir, f.dst)
	b.filesToCopy = append(b.filesToCopy, f)
	return f
}

func (b *ImageBuilder) AddGZippedBlob(blobPath string) OutputFile {
	f := OutputFile{
		src: blobPath,
		dst: filepath.Join(b.blobsDir, filepath.Base(blobPath)+".tar.gz"),
	}
	f.rel, _ = filepath.Rel(b.stagingDir, f.dst)
	b.filesToCopy = append(b.filesToCopy, f)
	return f
}

func NewImageBuilder(imageSha string, repoTags []string) ImageBuilder {
	builder := ImageBuilder{
		stagingDir: "/tmp/" + strings.Replace(imageSha, "sha256:", "", -1),
		repoTags:   repoTags,
	}
	builder.blobsDir = filepath.Join(builder.stagingDir, "blobs", "sha256")
	return builder
}
