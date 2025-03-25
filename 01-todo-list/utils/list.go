package utils

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

func ListTasks(f *os.File, verbose bool) {
	records, err := listRecords(f)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	defer w.Flush()

	if verbose {
		fmt.Fprintln(w, "\nID\tTASK\tCREATED AT\tCOMPLETED")
		fmt.Fprintln(w, "---\t----\t----------\t--------")
	} else {
		fmt.Fprintln(w, "\nID\tTASK\tCREATED AT")
		fmt.Fprintln(w, "---\t----\t----------")
	}
	
	for i, record := range records[1:] {
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
