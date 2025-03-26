package db

import (
	"encoding/csv"
	"fmt"
	"os"
)

const fileName = "db.csv"

func GetData() ([][]string, error) {
	// Open or create the CSV file
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening/creating file:", err)
		return nil, err
	}
	defer f.Close()
	// Reset file pointer to beginning
	f.Seek(0, 0)
	
	// Create a new CSV writer
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	if len(records) == 0 {
		writer := csv.NewWriter(f)
		// Write header
		headers := []string{"ID", "Task", "CreatedAt", "IsComplete"}
		writer.Write(headers)
		writer.Flush()
		records = [][]string{headers}
	}

	return records, nil
}

func WriteData(data [][]string){
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening/creating file:", err)
	}
	defer f.Close()

	if err := f.Truncate(0); err != nil {
		fmt.Println("Error truncating file:", err)
		return
	}

	f.Seek(0, 0)

	writer := csv.NewWriter(f)
	defer writer.Flush()

	writer.WriteAll(data)
}
