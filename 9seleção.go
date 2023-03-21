package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var num3 int
	fmt.Print("Informe um número:")
	fmt.Scan(&num1)
	fmt.Print("Informe um número:")
	fmt.Scan(&num2)
	fmt.Print("Informe um número:")
	fmt.Scan(&num3)
	if num1 < num2 < num3 {
		fmt.Print(num1, num2, num3)
	}
}
