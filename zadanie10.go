package main 

import (
	"fmt"
	"strings"
)

type text_stats struct {
	chars_count int
	words_count int
	sentences_count int
}

func texts(t string) * text_stats {
	var stats = new(text_stats)
	stats.chars_count = len(t)
	var words = strings.Fields(t)
	stats.words_count = len(words)

	var sentences strings.Count(t, ".") + strings.Count(t, "!") + strings.Count(t, "?") 
}
func main () {
	
}
//не дописано