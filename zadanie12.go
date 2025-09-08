package main

import (
	"fmt"
	"strconv"
)

const (
	BIN = iota
	DEC
	HEX
)

func ConvertNumber(value int64, target_base int) string {
	switch target_base {
	case BIN:
		return strconv.FormatInt(value, 2)
	case HEX:
		return strconv.FormatInt(value, 16)
	default:
		return strconv.Itoa(int(value))
	}
}

func main() {
	var number = int64(15)
	var binary_representation = ConvertNumber(number, BIN)
	var hexadecimal_representation = ConvertNumber(number, HEX)

	fmt.Println("Исходное число:", number)
	fmt.Println("2:", binary_representation)
	fmt.Println("16:", hexadecimal_representation)

}
