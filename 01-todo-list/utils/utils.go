package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func listRecords(f *os.File) ([][]string, error) {
	// Reset file pointer to beginning
	f.Seek(0, 0)
	
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return records, nil
}
	

func getLastID(records [][]string) (uint64, error) {
	if len(records) <= 2 { 
		return 0, nil
	}

	var maxID uint64 = 0

	for _, record := range records[1:] {
		if record[0] != "" {
			id, err := strconv.ParseUint(record[0], 10, 8)
			if err != nil {
				return 0, err
			}
			if id > maxID {
				maxID = id
			}
		}
	}
	
	return maxID, nil
}
	