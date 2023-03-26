package main

import (
	// "bufio"
	// "errors"
	"flag"
	"fmt"
	"github.com/heybran/todo-app"
	"log"
	// "io"
	// "os"
	// "strings"
)

const (
	todoFile = "/Users/brandon/codes/heybran/todos.json"
)

func main() {
	var add bool
	var task string
	var complete int
	var delete bool
	var list bool
	var cat string
	var update bool
	var id int

	flag.BoolVar(&add, "add", false, "add a new todo")
	flag.StringVar(&task, "task", "", "todo task content")
	flag.IntVar(&complete, "complete", 2, "mark a todo as completed")
	flag.BoolVar(&delete, "delete", false, "delete a todo")
	flag.BoolVar(&list, "list", false, "list all todos")
	flag.StringVar(&cat, "cat", "", "set todo category")
	flag.BoolVar(&update, "update", false, "update todo task")
	flag.IntVar(&id, "id", 0, "todo id to delete")

	flag.Parse()

	// https://pkg.go.dev/flag#hdr-Usage
	// will print: %!t(*bool=0xc0001a2002)
	// fmt.Printf("%t\n", add)

	// go run cmd/main.go -add
	// go run cmd/main.go --add
	// will all print true
	// fmt.Printf("%t\n", *add)

	fmt.Printf("Cat: %s\n", cat)
	fmt.Printf("Complete: %d\n", complete)
	fmt.Printf("Delete: %t\n", delete)
	fmt.Printf("Add: %t\n", add)
	fmt.Printf("Task: %s\n", task)

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		log.Fatal(err)
	}

	switch {
	case add:
		// if run: go run cmd/main.go -add hello world
		// will print: [hello world]
		// fmt.Print(flag.Args());
		// task, err := getInput(os.Stdin, flag.Args()...)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		todos.Add(task, cat)
		err := todos.Store(todoFile)
		if err != nil {
			log.Fatal(err)
		}
		todos.Print()

	case update:
		err := todos.Update(id, task, cat, complete)
		if err != nil {
			log.Fatal(err)
		}
		err = todos.Store(todoFile)
		if err != nil {
			log.Fatal(err)
		}
		todos.Print()

	case list:
		todos.Print()

	case delete:
		fmt.Print(delete)
		err := todos.Delete(id)
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
