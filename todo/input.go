package todo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func ReadATodo() Todo {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nEnter title:")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("\nEnter description:")
	desc, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return Todo{
		Title:     strings.Trim(title, "\n"),
		Desc:      strings.Trim(desc, "\n"),
		Done:      false,
		CreatedAt: time.Now(),
	}
}

func InputLoop(existingTodos []Todo) []Todo {
	todos := make([]Todo, 0)
	for {
		fmt.Printf("\nPlease choose an option\n")
		fmt.Printf("1. Create a new todo\n")
		fmt.Printf("2. Select a new todo\n")
		fmt.Printf("0. Exit the loop \n")
		fmt.Printf("Enter your choice: \n")
		var input string
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Printf("E#1VUOMM - %v\n", err)
		}

		if input == "0" {
			break
		} else if input == "1" {
			t := ReadATodo()
			t.Print()
			todos = append(todos, t)
		} else if input == "2" {
			for i, t := range existingTodos {
				t.PrintIdAndTitle(i + 1)
			}
			var selectTodoNumber int
			fmt.Printf("Please enter ID of todo to select it: ")
			_, err := fmt.Scanf("%d", &selectTodoNumber)
			if err != nil {
				fmt.Printf("E#1VUQF1 - %v\n", err)
			}
			selectTodoIndex := selectTodoNumber - 1
			option := SingleTodoActions(existingTodos[selectTodoIndex])
			switch option {
			case "1":
				fmt.Printf("\nYou want to mark the todo %v as done!", existingTodos[selectTodoIndex])
			case "2":
				fmt.Printf("\nYou want to edit the title of the Todo: %v", existingTodos[selectTodoIndex])
			default:
				fmt.Printf("\nInvalid Selection!")
			}
		} else {
			fmt.Println("You've selected an invalid option. Try again")
		}
	}
	return todos
}

func SingleTodoActions(todo Todo) string {
	fmt.Printf("\n You have selected: %v", todo.Title)
	fmt.Printf("\n 1. Mark as done")
	fmt.Printf("\n 2. Edit the title ")
	fmt.Printf("\n Please select the option: ")
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Printf("Fatal error: ", err)
	}
	return input
}
