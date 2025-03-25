package utils

import (
	"testing"
)


func TestGetLastID(t *testing.T) {
	t.Run("getting last id for empty records", func(t *testing.T) {
		records:= [][]string{}
		lastID, _ := getLastID(records)
		if lastID != 0 {
			t.Fatalf("Last id is not 0: %d", lastID)
		}
	})
	
	t.Run("getting last id for headers only", func(t *testing.T) {
		records:= [][]string{
			{"id", "task", "created_at", "completed"},
		}
		lastID, _ := getLastID(records)
		if lastID != 0 {
			t.Fatalf("Last id is not 0: %d", lastID)
		}
	})

	t.Run("getting last id for existing records", func(t *testing.T) {
		records:= [][]string{
			{"1", "test task", "2021-01-01T00:00:00Z", "false"},
			{"2", "test task 2", "2021-01-01T00:00:00Z", "false"},
		}

		lastID, _ := getLastID(records)
		
		if lastID != 2 {
			t.Fatalf("Last id is not 2: %d", lastID)
		}
	})

	t.Run("getting last id for unsorted records", func(t *testing.T) {
		records:= [][]string{
			{"2", "test task 2", "2021-01-01T00:00:00Z", "false"},
			{"1", "test task", "2021-01-01T00:00:00Z", "false"},
		}

		lastID, _ := getLastID(records)
		
		if lastID != 2 {
			t.Fatalf("Last id is not 2: %d", lastID)
		}
	})
}