package main

import (
	"fmt"
	"strings"
	"time"
)

const ADD, DELETE, MOVE, EXIT, HELP string = "add", "delete", "move", "exit", "help"

/*
TODO: INITIALIZE ToDoList with any saved dates
*/
func initializeToDoList() []string {

	l := make([]string, 0)
	return l
}

func printList(l *[]string) {

	list := *l
	date := time.Now().Format("Monday January 1, 2006")
	fmt.Printf("Today's date is: %s\n", date)
	if len(list) == 0 {
		fmt.Println("There are no items in your to do list.")
	} else {
		fmt.Println("To Do List:")
		for i, item := range list {
			fmt.Printf("%d. %s\n", (i + 1), item)
		}

	}
}

func main() {

	list := initializeToDoList()

	loop := true
	for loop {
		printList(&list)
		fmt.Println("Please input a command. Type help to see available commands.")
		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		switch input {

		case ADD:
			fmt.Print("Append: ")
			fmt.Scan(&input)
			list = append(list, input)
		case DELETE:
			fmt.Print("Delete: ")
		case MOVE:
			fmt.Print("Move: ")
		case EXIT:
			loop = false
		default:
			fmt.Println("Available commands: add, delete, move, exit, help")

		}
	}
}
