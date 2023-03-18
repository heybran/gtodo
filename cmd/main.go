package main

import (
	"bufio"
	"errors"
	"flag"
	// "fmt"
	"log"
	"github.com/heybran/todo-app"
	"io"
	"os"
	"strings"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	delete := flag.Int("delete", 0, "delete a todo")
	list := flag.Bool("list", false, "list all todos")

	flag.Parse()

	// https://pkg.go.dev/flag#hdr-Usage
	// will print: %!t(*bool=0xc0001a2002)
	// fmt.Printf("%t\n", add)

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
		// if run: go run cmd/main.go -add hello world
		// will print: [hello world]
		// fmt.Print(flag.Args());
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			log.Fatal(err)
		}

		todos.Add(task)
		err = todos.Store(todoFile)
		if err != nil {
			log.Fatal(err)
		}
		todos.Print()

	case *list:
		todos.Print()
	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			log.Fatal(err)
		}

		err = todos.Store(todoFile)
		if err != nil {
			log.Fatal(err)
		}
	case *delete > 0: 
		err := todos.Delete(*delete)
		if err != nil {
			log.Fatal(err)		
		}
		err = todos.Store(todoFile)
		if err != nil {
			log.Fatal(err)		
		}
		todos.Print()
	default:
		log.Fatal("invalid command")
	}
}


// getInput will get the input task typed in terminal
func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	// if len(args) means we type in: go run cmd/main.go -add
	// then NewScanner returns a new Scanner to read from user input
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("You didn't type in any task though...")
	}

	return text, nil
}