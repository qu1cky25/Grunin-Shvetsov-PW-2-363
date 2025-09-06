package main 

import (
	"fmt"
)

func count_votes (votes []string) {
	var candidate = map[string]int {"Anna": 0, "Boris": 0, "Victor": 0}
	var total_vote = len(votes)
	for i:=0; i<len(votes); i++ {
		var vote = votes[i]
		if _, f := candidate[vote]; f {
			candidate[vote]++
		}
	}

	for cand, count := range candidate {
		var p = (float64(count) / float64(total_vote)) * 100
		fmt.Println("Кандидат ", cand,  " получил ", count, " голосов. В процентах: ", p)
	}
}
func main () {
	var votes = []string {"Anna", "Boris", "Victor", "Boris", "Boris","Boris", "Boris", "Boris"}
	count_votes(votes)
}