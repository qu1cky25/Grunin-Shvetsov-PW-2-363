package main 

import (
	"fmt"
)

type log_entry struct {
	IP string
	response_code int
	times_tamp string 
}
func error_log (log_entries []log_entry) []log_entry {
	var logs = make([]log_entry, 0)
	for _, entry := range log_entries {
		if entry.response_code <= 599 && entry.response_code >= 400 {
			logs = append(logs, entry)
		}
	}
	return logs
}
func main () {
	var logs = []log_entry {
		{"192.168.1.1", 256 ,"2025.09.05-12:01:00"},
		{"192.168.1.2", 257 ,"2025.09.05-12:02:00"},
		{"192.168.1.3", 258 ,"2025.09.05-12:03:00"},
		{"192.168.1.4", 400 ,"2025.09.05-12:04:00"},
		{"192.168.1.5", 456 ,"2025.09.05-12:05:00"},
		{"192.168.1.6", 602 , "2025.09.05-12:06:00"},
	}
	var errors = error_log(logs)
	fmt.Println(errors)
}