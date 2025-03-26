package utils

import (
	"fmt"
	"todo-list/db"
)

func DeleteTask(data [][]string, taskID string) {
	fmt.Printf("Deleting task with ID: %s\n", taskID)

	if len(data) <= 1 {
		fmt.Println("No tasks to delete")
		return;
	}
	var filtered_records = filter(
		data[1:],
		func(row []string) bool {
			return row[0] != taskID
		},
	)

	var found = (len(data) - 1) != len(filtered_records)

	if found {
		result := append(data[:1], filtered_records...)
		db.WriteData(result)
		
		fmt.Println("Task deleted successfully")
	} else {
		fmt.Printf("Task with ID: %v was not found\n", taskID)
	}
}
