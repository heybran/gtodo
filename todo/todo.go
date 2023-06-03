package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexeyco/simpletable"
	"log"
	"os"
	"time"
	"strings"
)

// todo struct
type item struct {
	Task        string
	Category    string
	Done        bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// []item - slice
type Todos []item

// Add will add a new task to slice Todos
func (t *Todos) Add(task string, cat string) {
	todo := item{
		Task:        task,
		Category:    cat,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: nil, // set to nil instead of time.Time{}
	}

	// add a new task to Todos slice
	*t = append(*t, todo)
}

// Complete will mark requested task as completed
// not needed as this functionality is handled by Update() func
// func (t *Todos) Complete(index int) error {
// 	ls := *t
// 	if index <= 0 || index > len(ls) {
// 		return errors.New("invalid index")
// 	}

// 	ls[index-1].Done = true
// 	ls[index-1].CompletedAt = time.Now()

// 	return nil
// }

// Update will update task related to specific id
func (t *Todos) Update(index int, task string, cat string, done int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	if len(task) != 0 {
		ls[index-1].Task = task
	}

	if len(cat) != 0 {
		ls[index-1].Category = cat
	}

	// The error occurs because we are trying to assign a value of type time.Time
	// to a variable of type *time.Time.
	// To fix this, we can use the new() function to create a new pointer to time.Time
	// and assign it to CompletedAt.
	// By using &completedAt, we create a new pointer to time.Time that points to the completedAt variable
	if done == 0 {
		ls[index-1].Done = false
		ls[index-1].CompletedAt = nil
	} else if done == 1 {
		ls[index-1].Done = true
		completedAt := time.Now()
		ls[index-1].CompletedAt = &completedAt // create a new pointer to time.Time
	}

	return nil
}

// Delete will delete requested task from slice Todos
func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

// Load will read .todos.json file and update data into Todos slice
func (t *Todos) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	if len(data) == 0 {
		return err
	}

	err = json.Unmarshal(data, t)
	if err != nil {
		return err
	}

	return nil
}

// Store will write Todos slice data into .todos.json file
func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// Print will print out the current todo tasks
func (t *Todos) Print(status int, cat string) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Category"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	requestedTodos := []item{}

	for _, todo := range *t {
		if status == 1 {
			if todo.Done {
				requestedTodos = append(requestedTodos, todo)
			}
		}

		if status == 0 {
			if !todo.Done {
				requestedTodos = append(requestedTodos, todo)
			}
		}

		if status != 1 && status != 0 {
			requestedTodos = append(requestedTodos, todo)
		}
	}

	requestedCatTodos := []item{}

	for _, todo := range requestedTodos {
		if strings.ToLower(todo.Category) == strings.ToLower(cat) || cat == "" {
			requestedCatTodos = append(requestedCatTodos, todo)
		}
	}

	for i, item := range requestedCatTodos {
		i++
		task := item.Task
		done := "No"
		completedAt := ""

		if item.Done {
			task = fmt.Sprintf(item.Task)
			done = "\u2705"
		}

		if item.CompletedAt != nil {
			completedAt = item.CompletedAt.Format("2006-01-02")
		}

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", i)},
			{Text: item.Category},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format("2006-01-02")},
			// {Text: item.CompletedAt.Format("2006-01-02")},
			{Text: completedAt},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: ""},
		{Align: simpletable.AlignLeft, Span: 5, Text: fmt.Sprintf("You have %d pending todos", t.CountPending())},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

// CountPending() will print out the pending tasks
func (t *Todos) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}
