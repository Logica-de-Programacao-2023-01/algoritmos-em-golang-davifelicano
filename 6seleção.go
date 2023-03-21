package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	fmt.Print("Escreva um número:")
	fmt.Scan(&num1)
	fmt.Print("Escreva outro número:")
	fmt.Scan(&num2)
	x := num1 * num2
	y := num1 + num2
	if num1 > 0 && num2 > 0 {
		fmt.Print("Ambos são positivos então a multiplicação é:", x)
	} else if num1 < 0 && num2 > 0 {
		fmt.Print("Primeiro termo é negativo então a soma entre eles é:", y)
	} else if num1 > 0 && num2 < 0 {
		fmt.Print("Segudno termo é negativo então a sona entre eles é :", y)
	} else {
		fmt.Print("ambos são negativos então a soma entre eles é : ", y)
	}
}
