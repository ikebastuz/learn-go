package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func AddTask(f *os.File, task string) {
	fmt.Printf("Adding task: \"%v\"\n", task)
	
	records, err := listRecords(f)
	if err != nil {
		fmt.Println("Error getting records:", err)
		return
	}

	lastID, err := getLastID(records)
	if err != nil {
		fmt.Println("Error getting last ID:", err)
		return
	}

	f.Seek(0, 2)
	writer := csv.NewWriter(f)
	defer writer.Flush()

	record := []string{
		fmt.Sprintf("%d", lastID + 1),
		task,
		time.Now().Format(time.RFC3339),
		fmt.Sprintf("%v", false),
	}
	writer.Write(record)

	fmt.Println("Task added successfully")
}