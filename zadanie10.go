package main

import (
	"fmt"
	"strings"
)

type TextStats struct {
	chars_count     int
	words_count     int
	sentences_count int
}

func text(t string) *TextStats {
	var stats = new(TextStats)
	stats.chars_count = len(t)
	var words = strings.Fields(t)
	stats.words_count = len(words)
	var sentences = strings.Count(t, ".") +
		strings.Count(t, "!") +
		strings.Count(t, "?")
	stats.sentences_count = sentences
	return stats
}

func main() {
	var t = "Это простой тест. Сколько тут предложений?"
	var stats = text(t)
	fmt.Println("Кол-во символов:", stats.chars_count, ". Кол-во слов:", stats.words_count, ". Кол-во предложений:", stats.sentences_count, ".")
}
