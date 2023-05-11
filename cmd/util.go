package cmd

import (
	"os"
	"log"
	"path/filepath"
)

// getJsonFile will grab the .todos.json file located at user home directory
func GetJsonFile() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(homeDir, ".todos.json")
}
