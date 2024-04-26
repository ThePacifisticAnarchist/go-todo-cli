package main

import (
	"encoding/json"
	"fmt"
	"github.com/ThePacifisticAnarchist/todo-cli/todo"
	"time"
)

func main() {
	fmt.Printf("This is a CLI app to manage todo lists. Its written in golang\n")
	t := todo.Todo{
		Title:     "",
		Desc:      "Just checking app",
		Done:      false,
		CreatedAt: time.Now(),
	}

	//t.Print()
	if t.Validate() != nil {
		fmt.Printf("Todo is not valid. Error: %v\n", t.Validate())
	} else {
		fmt.Printf("The todo is valid\n")
	}

	todoBytes, err := json.Marshal(t)
	if err != nil {
		fmt.Printf("Something went wrong while encoding to json: %v\n", err)
	}
	fmt.Printf("\nJSON Bytes: %v\n", todoBytes)
	todoJsonString := string(todoBytes)
	fmt.Printf("\nJSON String: %v\n", string(todoJsonString))
}
