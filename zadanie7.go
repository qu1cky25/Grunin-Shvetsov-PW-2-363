package main 

import (
	"fmt"
)

type Employee struct {
	ID	int
	Name	string
	Position	string
	Salary	float64
} 

func calculation_salary(employee []Employee) (float64, float64) {
	var total_salary float64
	for _, employ := range employee {
		total_salary += employ.Salary
	}
	var average_salary float64 = total_salary / float64(len(employee))
	return total_salary, average_salary 
}

func main () {
	employee := []Employee{
		{ID: 121, Name: "Alex Mirror", Position: "Director", Salary: 120000},
		{ID: 122, Name: "Max Verstappen", Position: "Driver", Salary:80000 },
		{ID: 121, Name: "Kato Repper", Position: "Cleaner", Salary: 50000},
	}	

	var totalsalary, averagesalary = calculation_salary(employee)
	fmt.Println("Общий фонд зарплаты: ", totalsalary, ". Средняя зарплата: ", averagesalary, ".")
}