package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// WriteJSONToFile writes a struct to a file in JSON format
func WriteJSONToFile(v interface{}, filename string) error {
	// Marshal the struct to JSON
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	// Write JSON data to a file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	return err
}

func loadJSON(path string, output any) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("failed to open index file: %w", err)
	}

	if err := json.Unmarshal(byteValue, output); err != nil {
		return fmt.Errorf("malformed file: %w", err)
	}
	return nil
}
