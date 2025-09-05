package main 

import (
	"fmt"
)

func main () {
	var VesBag float64
	fmt.Println("Введите вес основного багажа")
	fmt.Scanln(&VesBag)

	var HandBag float64
	fmt.Println("Введите вес ручной клади")
	fmt.Scanln(&HandBag)

	var AddHandBag float64
	fmt.Println("Введите вес допольнительной ручной клади")
	fmt.Scanln(&AddHandBag)

	var totalWeight = VesBag + HandBag + AddHandBag 
	fmt.Println("Общий вес: ",totalWeight)

}