package main

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

func addFileToTar(tw *tar.Writer, baseDir, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	// Ensure the header name is relative to baseDir
	relPath, err := filepath.Rel(baseDir, path)
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(stat, relPath)
	if err != nil {
		return err
	}

	header.Name = relPath

	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	if !stat.IsDir() {
		_, err = io.Copy(tw, file)
	}

	return err
}

func addDirToTar(tw *tar.Writer, baseDir, path string) error {
	return filepath.Walk(path, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if file == path {
			// Skip the root directory itself
			return nil
		}

		return addFileToTar(tw, baseDir, file)
	})
}

func createTar(baseDir, tarName string, files []string) error {
	tarFile, err := os.Create(tarName)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	tw := tar.NewWriter(tarFile)
	defer tw.Close()

	for _, file := range files {
		fullPath := filepath.Join(baseDir, file)
		fi, err := os.Stat(fullPath)
		if err != nil {
			return err
		}

		if fi.IsDir() {
			err = addDirToTar(tw, baseDir, fullPath)
		} else {
			err = addFileToTar(tw, baseDir, fullPath)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
