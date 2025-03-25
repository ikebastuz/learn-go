package main

import (
	"fmt"
	"os"
	"todo-list/db"
	"todo-list/utils"
)

func main() {
	file, _ := db.InitDB()
	defer file.Close()

	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
			case "list":
				verbose := len(args) > 1 && (args[1] == "-a" || args[1] == "--all")
				utils.ListTasks(file, verbose)
				return
			case "add":
				if len(args) < 2 {
					fmt.Println("Please provide a task to add")
					return
				}
				utils.AddTask(file, args[1])
				return
			case "complete":
				if len(args) < 2 {
					fmt.Println("Please provide a task ID to complete")
					return
				}
				utils.CompleteTask(file, args[1])
				return
			default:
				fmt.Println("Invalid command")
		}
	}
}
