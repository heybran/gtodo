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
	// 定义"add"子命令来添加待办事项
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	// Define an optional "--cat" flag for the todo item
	todoCat := addCmd.String("cat", "Uncategorized", "The category of the todo item")

	// listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	todoID := deleteCmd.Int("id", 0, "The id of todo to be delete")
	
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

	// Parse the command line arguments
	// 解析命令行参数
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

	// Check which subcommand was invoked
	// 检查调用了哪个子命令
	switch flag.Arg(0) {
	case "add":
		// Parse the arguments for the "add" subcommand
		// 解析"add"子命令的参数
		fmt.Printf("os args: %s\n", os.Args)
		addCmd.Parse(os.Args[2:])

		// Check if the required todo text was provided
		// 检查是否提供了必需的待办事项文本
		if addCmd.NArg() == 0 {
			fmt.Println("Error: the todo text is required for the 'add' subcommand")
			os.Exit(1)		
		}

		// Get the todo text from the positional argument
		// 从位置参数中获取待办事项文本
		todoText := addCmd.Arg(0)
		// fmt.Printf("%s\n", todoText)
		// 将待办事项和类别添加到Todos slice中
		todos.Add(todoText, *todoCat)
		// 将Todos slice写入todoFile文件中
		err := todos.Store(todoFile)
		if err != nil {
			log.Fatal(err)
		}
		todos.Print()
		fmt.Println("Todo item added successfully")
	case "list":
		todos.Print()
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		// if deleteCmd.NArg() == 0 {
		// 	fmt.Println("Error: the todo id is required for the 'delete' subcommand")
		// 	os.Exit(1)
		// }

		err := todos.Delete(*todoID)
		if err != nil {
			log.Fatal(err)
		}

		err = todos.Store(todoFile)
		if err != nil {
			log.Fatal(err)
		}

		todos.Print()
		
	default:
		fmt.Println("Error: invalid subcommand")
		os.Exit(1)
	}

	return

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
