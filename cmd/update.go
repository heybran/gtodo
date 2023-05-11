// Package cmd provides command-line subcommands for the gtodo application.
package cmd

import (
	"github.com/heybran/gtodo/todo"
	"flag"
	"fmt"
	"log"
	"os"
)

func UpdateTask(todos *todo.Todos, args []string) {	
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateID := updateCmd.Int("id", 0, "The id of todo to be updated")
	updateCat := updateCmd.String("cat", "", "The to-be-updated category of todo")
	updateTask := updateCmd.String("task", "", "To to-be-updated content of todo")
	updateDone := updateCmd.Int("done", 2, "The to-be-updated status of todo")

	// Parse the arguments for the "update" subcommand
	updateCmd.Parse(args)

	if *updateID == 0 {
		fmt.Println("Error: the --id flag is required for the 'update' subcommand.")
		os.Exit(1)		
	}
	err := todos.Update(*updateID, *updateTask, *updateCat, *updateDone)
	if err != nil {
		log.Fatal(err)		
	}

	err = todos.Store(GetJsonFile())
	if err != nil {
		log.Fatal(err)
	}

	todos.Print(2)
	fmt.Println("Todo item updated successfully.")
}