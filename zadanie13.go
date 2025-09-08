package main

import (
	"fmt"
)

func main() {
	expenses := map[string]float64{
		"Еда":         15000,
		"Транспорт":   5000,
		"Развлечения": 3000,
	}

	expenses["Еда"] += 2000

	total_expenses := 0.0
	for _, amount := range expenses {
		total_expenses += amount
	}

	fmt.Println("Расходы по категориям:")
	for category, amount := range expenses {
		fmt.Printf("%s: %.2f руб.\\n", category, amount)
	}

	fmt.Println("Общая сумма расходов:", total_expenses)
}
