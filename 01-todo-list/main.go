package main

import (
	"fmt"
	"os"
	"todo-list/db"
	"todo-list/utils"
)

func main() {
	data, _ := db.GetData()

	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
			case "list":
				verbose := len(args) > 1 && (args[1] == "-a" || args[1] == "--all")
				utils.ListTasks(data, verbose)
				return
			case "add":
				if len(args) < 2 {
					fmt.Println("Please provide a task to add")
					return
				}
				utils.AddTask(data, args[1])
				return
			case "complete":
				if len(args) < 2 {
					fmt.Println("Please provide a task ID to complete")
					return
				}
				utils.CompleteTask(data, args[1])
				return
			case "delete": {}
				if len(args) < 2 {
					fmt.Println("Please provide a task ID to delete")
					return
				}
				utils.DeleteTask(data, args[1])
				return
			default:
				fmt.Println("Invalid command")
		}
	}
}
