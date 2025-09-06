package main 

import (
	"fmt"
)

func collect_tags (posts [][]string) map[string]bool {
	var set = make(map[string]bool)
	for i :=0; i<len(posts); i++ {
		var current = posts[i]
		for j := 0; j < len(current); j++ {
			var tag = current[j]
			set[tag] = true
		}
	}
	return set 
}

func main () {
	 posts := [][]string{
		{"go", "backend"},
		{"git", "go", "tools"},
	}
	var unique_tags = collect_tags(posts)
	fmt.Println("Уникальные теги: ", unique_tags)
}