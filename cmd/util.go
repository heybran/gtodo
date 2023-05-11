package cmd

import (
	"os"
	"github.com/heybran/gtodo/todo"
	"log"
	"path/filepath"
	"strings"
	"bufio"
	"fmt"
)

// getJsonFile will grab the .todos.json file located at user home directory
func GetJsonFile() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(homeDir, ".todos.json")
}

// getUserApproval will get the user's approval when creating an empty json file
func GetUserApproval() (bool) {
	confirmMessage := "Need to create an empty \".todos.json\" file in your home directory to store your todo items, continue? (y/n): "

	r := bufio.NewReader(os.Stdin)
	var s string

	fmt.Print(confirmMessage)
	s, _ = r.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	for {
		if s == "y" || s == "yes" {
			return true
		}
		if s == "n" || s == "no" {
			return false
		}
	}
}

func RemindInit(todos *todo.Todos) {
	// check if .todos.json already exists in user home directory
	_, err := os.Stat(GetJsonFile())
	if err != nil {
		fmt.Println("Please run \"init\" subcommand to create an JSON file to store your todo items.")
		os.Exit(1)
	} else {
		if err := todos.Load(GetJsonFile()); err != nil {
			log.Fatal(err)
		}
	}
}