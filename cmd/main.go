package main

import (
	// "bufio"
	// "errors"
	"flag"
	"fmt"
	"github.com/heybran/todo-app"
	"log"
	// "io"
	"os"
	// "strings"
)

const (
	todoFile = "/Users/brandon/codes/heybran/todos.json"
)

func main() {
	// Define the "add" subcommand to add todo item
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTask := addCmd.String("task", "", "The content of new todo item")
	// Define an optional "--cat" flag for the todo item
	addCat := addCmd.String("cat", "Uncategorized", "The category of the todo item")

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

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listDone := listCmd.Int("done", 2, "The status of todo to be printed")
	
	// Parse the command line arguments
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		log.Fatal(err)
	}

	// Check which subcommand was invoked
	switch flag.Arg(0) {
	case "add":
		// Parse the arguments for the "add" subcommand
		addCmd.Parse(os.Args[2:])

		// Check if the required todo text was provided

		// if we're going with this route: todo add --task="Hello World" --cat="Hi"
		// then addCmd.NArg() will be 0
		// if addCmd.NArg() == 0 {
		// 	fmt.Println("Error: the todo text is required for the 'add' subcommand")
		// 	os.Exit(1)		
		// }

		if len(*addTask) == 0 {
			fmt.Println("Error: the --task flag is required for the 'add' subcommand")
			os.Exit(1)		
		}

		// Get the todo text from the positional argument
		todos.Add(*addTask, *addCat)
		err := todos.Store(todoFile)
		if err != nil {
			log.Fatal(err)
		}

		todos.Print(2)
		fmt.Println("Todo item added successfully")
	case "list":
		// Parse the arguments for the "list" subcommand
		listCmd.Parse(os.Args[2:])
		todos.Print(*listDone)
	case "delete":
		deleteCmd.Parse(os.Args[2:])

		err := todos.Delete(*deleteID)
		if err != nil {
			log.Fatal(err)
		}

		err = todos.Store(todoFile)
		if err != nil {
			log.Fatal(err)
		}

		todos.Print(2)
		fmt.Println("Todo item deleted successfully")
	case "update":
		updateCmd.Parse(os.Args[2:])

		if *updateID == 0 {
			fmt.Println("Error: the --id flag is required for the 'update' subcommand")
			os.Exit(1)		
		}
		err := todos.Update(*updateID, *updateTask, *updateCat, *updateDone)
		if err != nil {
			log.Fatal(err)		
		}

		err = todos.Store(todoFile)
		if err != nil {
			log.Fatal(err)
		}

		todos.Print(2)
		fmt.Println("Todo item updated successfully")
	default:
		fmt.Println("Error: invalid subcommand")
		os.Exit(1)
	}
}

// getInput will get the input task typed in terminal
// func getInput(r io.Reader, args ...string) (string, error) {
// 	if len(args) > 0 {
// 		return strings.Join(args, " "), nil
// 	}

// 	// if len(args) means we type in: go run cmd/main.go -add
// 	// then NewScanner returns a new Scanner to read from user input
// 	scanner := bufio.NewScanner(r)
// 	scanner.Scan()
// 	if err := scanner.Err(); err != nil {
// 		return "", err
// 	}

// 	text := scanner.Text()

// 	if len(text) == 0 {
// 		return "", errors.New("You didn't type in any task though...")
// 	}

// 	return text, nil
// }
