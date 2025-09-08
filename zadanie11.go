package main

import (
	"fmt"
)

type Product struct {
	name     string
	category string
	price    float64
}

func FilterProducts(products []Product, max_price float64, category string) []Product {
	var filtered_products []Product
	for _, p := range products {
		if p.price <= max_price && p.category == category {
			filtered_products = append(filtered_products, p)
		}
	}
	return filtered_products
}

func main() {

	products := []Product{
		{"Яблоко", "Фрукты", 50},    
		{"Картофель", "Овощи", 80},  
		{"Апельсин", "Фрукты", 100}, 
		{"Огурец", "Овощи", 40},     
		{"Манго", "Фрукты", 60},     
	}

	filtered := FilterProducts(products, 90, "Фрукты")

	fmt.Println("Отфильтрованные товары:")
	for _, prod := range filtered {
		fmt.Printf("- Название: %s, Категория: %s, Цена: %.2f руб.\\n",
			prod.name, prod.category, prod.price)
	}

}
