package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CompleteTask(f *os.File, taskID string) {
	fmt.Printf("Completing task with ID: %s\n", taskID)
	
	// Reset file pointer to beginning
	f.Seek(0, 0)
	
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Truncate the file to remove all existing content
	if err := f.Truncate(0); err != nil {
		fmt.Println("Error truncating file:", err)
		return
	}

	// Reset file pointer to beginning after truncate
	f.Seek(0, 0)
	
	writer := csv.NewWriter(f)
	defer writer.Flush()

	// Write all records, updating the specified task
	for _, record := range records {
		if len(record) >= 4 {
			if record[0] == taskID {
				record[3] = "true"
			}
			if err := writer.Write(record); err != nil {
				fmt.Println("Error writing record:", err)
				return
			}
		}
	}

	fmt.Println("Task completed successfully")
}
