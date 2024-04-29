package main

import (
	"encoding/json"
	"fmt"
	"github.com/ThePacifisticAnarchist/todo-cli/todo"
	"os"
)

func main() {
	fmt.Printf("This is a CLI app to manage todo lists. Its written in golang\n")

	f, err := os.OpenFile("/home/ganesh/GolandProjects/workshop/ToDo/todos.json", os.O_CREATE|
		os.O_RDWR, 0644)

	if err != nil {
		fmt.Printf("Error when opening file: %v\n", err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("Some fatal error: %v", err)
		}
	}(f)

	// Read the contents of the existing file...
	dat, err := os.ReadFile("/home/ganesh/GolandProjects/workshop/ToDo/todos.json")
	// fmt.Print(string(dat))

	var existingTodos []todo.Todo

	// ... and try to cast it in a slice of todos
	err = json.Unmarshal(dat, &existingTodos)
	if err != nil {
		fmt.Printf("Fatal error when reading file: %v", err)
	}

	// ... and print them!
	for _, t := range existingTodos {
		t.Print()
		fmt.Printf("####")
	}

	todos := todo.InputLoop(existingTodos)
	for _, t := range todos {
		t.Print()
		fmt.Printf("####")
	}

	// Let's convert the list of Todos to JSON
	todoJsonBytes, err := json.MarshalIndent(todos, "", "")
	if err != nil {
		fmt.Printf("Error when marshalling to JSON: %v", err)
		return
	}

	todoJsonContents := string(todoJsonBytes)

	_, err = f.WriteString(todoJsonContents)
	if err != nil {
		fmt.Printf("Error when writing to file: %v\n", err)
	}
	fmt.Printf("Wrote contents to file")

}
