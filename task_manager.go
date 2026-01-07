package main

import (
	"fmt"
	"os"
)

var tasks []string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: add | list | delete <index>")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Task name required")
			return
		}
		tasks = append(tasks, os.Args[2])
		fmt.Println("Task added")

	case "list":
		if len(tasks) == 0 {
			fmt.Println("No tasks available")
			return
		}
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task)
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Task index required")
			return
		}
		fmt.Println("Delete feature placeholder")

	default:
		fmt.Println("Unknown command")
	}
}
