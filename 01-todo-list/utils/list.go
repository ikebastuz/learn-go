package utils

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

func ListTasks(data [][]string, verbose bool) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	defer w.Flush()

	if len(data) < 2 {
		fmt.Println("No tasks")
		return
	}

	to_show_records := filter(
		data[1:], 
		func(record []string) bool {
			if verbose {
				return true
			}
			return record[3] == "false"
		},
	)

	if len(to_show_records) == 0 {
		fmt.Println("No tasks")
		return
	}

	if verbose {
		fmt.Fprintln(w, "\nID\tTASK\tCREATED AT\tCOMPLETED")
		fmt.Fprintln(w, "---\t----\t----------\t--------")
	} else {
		fmt.Fprintln(w, "\nID\tTASK\tCREATED AT")
		fmt.Fprintln(w, "---\t----\t----------")
	}
	
	for i, record := range to_show_records {
		if len(record) >= 4 {
			id, err := strconv.ParseUint(record[0], 10, 8)
			if err != nil {
				fmt.Printf("Error parsing ID for task: %v\n", err)
				continue
			}
			
			createdAt, err := time.Parse(time.RFC3339, record[2])
			createdAtStr := timediff.TimeDiff(createdAt)
			if err != nil {
				fmt.Printf("Error parsing date for task %s: %v\n", record[0], err)
				continue
			}

			if record[3] == "true" && !verbose {
				continue
			}
			
			if verbose {
				fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", id, record[1], createdAtStr, record[3])
			} else {
				fmt.Fprintf(w, "%d\t%s\t%s\n", id, record[1], createdAtStr)
			}
		} else {
			fmt.Printf("Record %d has insufficient fields: %v\n", i, record)
		}
	}
}
