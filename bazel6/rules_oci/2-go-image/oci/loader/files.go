package main

import (
	"io"
	"os"
)

// CopyFile copies a file from src to dst. If dst does not exist, it will be created.
// If it exists, it will be overwritten.
func copyFile(src, dst string) error {
	// Open the source file
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create the destination file
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copy the contents
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// Commit the file contents
	err = destFile.Sync()
	return err
}
