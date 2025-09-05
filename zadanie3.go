package main 

import (
	"fmt"
)

type Order struct {
	id		int 
	items	[]int
	total	float64
	address	string 
	completed bool
}
var orders = make(map[int]*Order)

func addorder (id int, items []int, total float64, address string) {
	var order = &Order{id, items, total, address, false}
	orders[id] = order 
}

func main () {
	addorder(1, []int{1013, 1150, 1456}, 2000.0, "Курган, р-н Северный")
	for _, add := range orders{
		fmt.Println(add.id, add.items, add.total, add.address, add.completed)
	}

}