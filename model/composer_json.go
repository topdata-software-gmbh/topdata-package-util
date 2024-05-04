package model

import (
	"encoding/json"
	"io"
	"os"
)

// ComposerJSON represents the structure of the composer.json file
type ComposerJSON struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Version     string            `json:"version"`
	Require     map[string]string `json:"require"`
}

// LoadFromFile loads data from a composer.json file into a ComposerJSON struct
func (c *ComposerJSON) LoadFromFile(filePath string) error {
	// Read the content of composer.json file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the JSON content
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Unmarshal JSON data into struct
	err = json.Unmarshal(byteValue, c)
	if err != nil {
		return err
	}

	return nil
}
