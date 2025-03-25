package db

import (
	"encoding/csv"
	"fmt"
	"os"
)

const fileName = "db.csv"

func InitDB() (*os.File, error) {
	// Open or create the CSV file
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening/creating file:", err)
		return nil, err
	}
	// Reset file pointer to beginning
	f.Seek(0, 0)
	
	// Create a new CSV writer
	reader := csv.NewReader(f)
	writer := csv.NewWriter(f)
	defer writer.Flush()
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	if len(records) == 0 {
		// Write header
		header := []string{"ID", "Task", "CreatedAt", "IsComplete"}
		writer.Write(header)
	}

	return f, nil
}
