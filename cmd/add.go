// Package cmd provides command-line subcommands for the gtodo application.
package cmd

import (
	"github.com/heybran/gtodo/todo"
	"os"
	"flag"
	"fmt"
	"log"
)

func AddTask(todos *todo.Todos, args []string) {	
	// Define the "add" subcommand to add todo item
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTask := addCmd.String("task", "", "The content of new todo item")
	// Define an optional "--cat" flag for the todo item
	addCat := addCmd.String("cat", "Uncategorized", "The category of the todo item")

	// Parse the arguments for the "add" subcommand
	addCmd.Parse(args)

	// Check if the required todo text was provided

	// if we're going with this route: todo add --task="Hello World" --cat="Hi"
	// then addCmd.NArg() will be 0
	// if addCmd.NArg() == 0 {
	// 	fmt.Println("Error: the todo text is required for the 'add' subcommand")
	// 	os.Exit(1)		
	// }

	if len(*addTask) == 0 {
		fmt.Println("Error: the --task flag is required for the 'add' subcommand.")
		os.Exit(1)		
	}

	// Get the todo text from the positional argument
	todos.Add(*addTask, *addCat)
	err := todos.Store(GetJsonFile())
	if err != nil {
		log.Fatal(err)
	}

	todos.Print(2, "")
	fmt.Println("Todo item added successfully.")
}