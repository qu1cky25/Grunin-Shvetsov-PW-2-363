package main 

import (
	"fmt"
	"strings"
)

func validateUser(name string, age int, email string) error {
	if name == "" || len(name) > 50 {
		fmt.Println("Имя должно быть не больше 50 символов")
	} 
	if age<18 || age>120  {
		fmt.Println("Возраст должен быть от  18 лет и не больше 120")
	} 
	if strings.Index(email, "@") == -1 {
		fmt.Println("Почта должна содержать @")		
	}
	return nil

}
func main () {
	var name string
	fmt.Println("Введите свое имя!")
	fmt.Scan(&name)

	var age int
	fmt.Println("Введите свой возраст!")
	fmt.Scan(&age)

	var email string
	fmt.Println("Введите свою почту!")
	fmt.Scan(&email)

	validateUser(name, age, email)
	
}