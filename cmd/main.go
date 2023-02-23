package main

import (
	// "bufio"
	// "errors"
	"flag"
	"fmt"
	"log"
	// "github.com/heybran/todo-app"
	// "io"
	// "os"
	// "strings"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	// complete := flag.Int("complete", 0, "mark a todo as completed")
	// delete := flag.Int("delete", 0, "delete a todo")
	// list := flag.Bool("list", false, "list all todos")

	flag.Parse()

	// https://pkg.go.dev/flag#hdr-Usage
	// will print: %!t(*bool=0xc0001a2002)
	fmt.Printf("%t\n", add)

	// go run cmd/main.go -add
	// go run cmd/main.go --add
	// will all print true
	// fmt.Printf("%t\n", *add)
	
	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		log.Fatal(err)
	}

	switch {
	case *add:
				
	}
}