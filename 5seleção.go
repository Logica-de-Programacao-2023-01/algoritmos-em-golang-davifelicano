package main

import "fmt"

func main() {
	var num1 int
	fmt.Print("Informe um numero:")
	fmt.Scan(&num1)
	if num1%3 == 0 && num1%5 == 0 {
		fmt.Print(num1, "È multiplo de 3 e 5 ")
	} else {
		fmt.Print(num1, "Não é multiplo de 3 e 5")
	}
}
