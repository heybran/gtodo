package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexeyco/simpletable"
	"log"
	"os"
	"time"
)

// todo struct
type item struct {
	Task		string
	Done		bool
	CreatedAt 	time.Time
	CompletedAt time.Time
}

// []item - slice
type Todos []item

// Add will add a new task to slice Todos
func (t *Todos) Add(task string) {
	todo := item{
		Task:		 task
		Done: 		 false,
		CreatedAt:	 time.Now(),
		CompletedAt: time.Time{}		
	}

	// add a new task to Todos slice
	*t = append(*t, todo)
}

// Complete will mark requested task as completed
func (t *Todos) Complete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].Done = true
	ls[index-1].CompletedAt = time.Now()

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

// Load
// func (t *Todos) Load(filename string) error {
// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		return log.Fatal(err)		
// 	}

	
// }


