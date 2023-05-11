// Package cmd provides command-line subcommands for the gtodo application.
package cmd

import (
	"os"
	"fmt"
	"log"
	"path/filepath"
)

func Init() {	
	ok := GetUserApproval()
	if !ok {
		fmt.Print("You've declined to create the JSON file. You can always run \"init\" subcommand again if you change your mind.")	
		os.Exit(0)
	}
	
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	filepath := filepath.Join(homeDir, ".todos.json")
	// check if .todos.json already exists in user home directory
	_, err = os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(filepath)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			fmt.Println("Succefully create a \".todos.json\" file in your home directory.")
		} else {
			log.Fatal("Unknown error occurred.")
		}
	} else {
		fmt.Print(".todos.json file exists in your home directory already.")	
	}
}