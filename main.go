package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/heybran/gtodo/todo"
	"github.com/heybran/gtodo/cmd"
	"log"
	"os"
	"strings"
	"path/filepath"
)

func main() {
	// listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	// If no --id=1 flag defined, todoID will default to 0
	// but if --id is present but didn't set any value, an error will be thrown
	deleteID := deleteCmd.Int("id", 0, "The id of todo to be deleted")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateID := updateCmd.Int("id", 0, "The id of todo to be updated")
	updateCat := updateCmd.String("cat", "", "The to-be-updated category of todo")
	updateTask := updateCmd.String("task", "", "To to-be-updated content of todo")
	updateDone := updateCmd.Int("done", 2, "The to-be-updated status of todo")

	
	// Parse the command line arguments
	flag.Parse()

	todos := &todo.Todos{}

	// Check which subcommand was invoked
	switch flag.Arg(0) {
	case "init":
		ok := getUserApproval()
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
	case "add":
		remindInit(todos)
		cmd.AddTask(todos, os.Args[2:])
	case "list":
		remindInit(todos)
		cmd.ListTasks(todos, os.Args[2:])
	case "delete":
		remindInit(todos)
		deleteCmd.Parse(os.Args[2:])

		err := todos.Delete(*deleteID)
		if err != nil {
			log.Fatal(err)
		}

		err = todos.Store(getJsonFile())
		if err != nil {
			log.Fatal(err)
		}

		todos.Print(2)
		fmt.Println("Todo item deleted successfully.")
	case "update":
		remindInit(todos)
		updateCmd.Parse(os.Args[2:])

		if *updateID == 0 {
			fmt.Println("Error: the --id flag is required for the 'update' subcommand.")
			os.Exit(1)		
		}
		err := todos.Update(*updateID, *updateTask, *updateCat, *updateDone)
		if err != nil {
			log.Fatal(err)		
		}

		err = todos.Store(getJsonFile())
		if err != nil {
			log.Fatal(err)
		}

		todos.Print(2)
		fmt.Println("Todo item updated successfully.")
	default:
		fmt.Println("Error: invalid subcommand.")
		os.Exit(1)
	}
}

func remindInit(todos *todo.Todos) {
	// check if .todos.json already exists in user home directory
	_, err := os.Stat(getJsonFile())
	if err != nil {
		fmt.Println("Please run \"init\" subcommand to create an JSON file to store your todo items.")
		os.Exit(1)
	} else {
		if err := todos.Load(getJsonFile()); err != nil {
			log.Fatal(err)
		}
	}
}

// getJsonFile will grab the .todos.json file located at user home directory
func getJsonFile() (string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(homeDir, ".todos.json")
}

// getUserApproval will get the user's approval when creating an empty json file
func getUserApproval() (bool) {
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
