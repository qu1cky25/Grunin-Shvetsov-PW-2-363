package main 

import (
	"fmt"
)

func main () {
	var dates = []int{9, 10, 11, 12, 13, 14, 25, 26}
	var one_price = 2100
	var two_price = 2850

	var sum = 0
	for _, days := range dates {
		var day = (days-1) % 7
		if day < 4 {
			sum += one_price
		} else {
			sum+= two_price
		}
	} 
	fmt.Println("Сумма проживания: ", sum)
}