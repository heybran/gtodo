// Package cmd provides command-line subcommands for the gtodo application.
package cmd

import (
	"github.com/heybran/gtodo/todo"
	"flag"
)

func ListTasks(todos *todo.Todos, args []string) {	
	// Define the "list" subcommand to list todo items
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listDone := listCmd.Int("done", 2, "The status of todo to be printed")

	// Parse the arguments for the "list" subcommand
	listCmd.Parse(args)
	todos.Print(*listDone)
}