package main

import (
	"fmt"
)

type InventoryItem struct {
	name       string
	weight     float64
	quest_item bool
}

func total_weight(inventory_items []InventoryItem) float64 {
	var totalweight float64 = 0
	for _, i := range inventory_items {
		totalweight += i.weight
	}
	return totalweight
}

func main() {
	items := []InventoryItem{
		{name: "Лук и стрелы", weight: 1.5, quest_item: true},
		{name: "Шлем", weight: 1.0, quest_item: false},
		{name: "Зелье ярости", weight: 0.3, quest_item: false},
		{name: "Меч", weight: 2.0, quest_item: true},
		{name: "Зелье магии", weight: 0.3, quest_item: false},
	}

	var totalweight = total_weight(items)
	fmt.Println("Вес всего инвентаря: ", totalweight)
}
