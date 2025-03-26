package utils

import (
	"fmt"
	"todo-list/db"
)

func CompleteTask(data [][]string, taskID string) {
	fmt.Printf("Completing task with ID: %s\n", taskID)

	var found bool = false
	for _, record := range data[1:] {
		if len(record) >= 4 {
			if record[0] == taskID {
				record[3] = "true"
				found = true
			}
		}
	}

	db.WriteData(data)

	if found {
		fmt.Println("Task completed successfully")
	} else {
		fmt.Printf("Task with ID: %v was not found\n", taskID)
	}
	
}
