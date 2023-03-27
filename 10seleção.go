package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Informe a idade:")
	fmt.Scan(&idade)
	if idade <= 9 {
		fmt.Print("Mirim")
	} else if idade > 9 && idade < 14 {
		fmt.Print("Infantil")
	} else if idade > 13 && idade < 18 {
		fmt.Print("Juvenil")
	} else if idade > 18 {
		fmt.Print("Adulto")
	}
}
