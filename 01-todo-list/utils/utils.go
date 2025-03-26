package utils

import (
	"strconv"
)

func getLastID(records [][]string) (uint64, error) {
	if len(records) < 1 { 
		return 0, nil
	}

	var maxID uint64 = 0

	for _, record := range records {
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
	
func filter[T any](ss []T, test func(T) bool)(ret []T){
	for _, s := range ss {
		if test(s){
			ret = append(ret, s)
		}
	}
	return
}