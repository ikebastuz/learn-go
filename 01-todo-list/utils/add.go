package utils

import (
	"fmt"
	"time"
	"todo-list/db"
)

func AddTask(data [][]string, task string) {
	fmt.Printf("Adding task: \"%v\"\n", task)

	lastID, err := getLastID(data[1:])
	if err != nil {
		fmt.Println("Error getting last ID:", err)
		return
	}

	new_record := []string{
		fmt.Sprintf("%d", lastID + 1),
		task,
		time.Now().Format(time.RFC3339),
		fmt.Sprintf("%v", false),
	}
	data = append(data, new_record)


	db.WriteData(data)
}