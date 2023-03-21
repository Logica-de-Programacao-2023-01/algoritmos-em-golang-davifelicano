package main

import "fmt"

func main() {
	var num1 int
	fmt.Print("Informe um número de 1 a 7 : ")
	fmt.Scan(&num1)
	if num1 == 1 {
		fmt.Print("Domingo")
	} else if num1 == 2 {
		fmt.Print("Segunda-Feira")
	} else if num1 == 3 {
		fmt.Print("terça-Feira")
	} else if num1 == 4 {
		fmt.Print("Quarta-Feira")
	} else if num1 == 5 {
		fmt.Print("Quinta-Feira")
	} else if num1 == 6 {
		fmt.Print("Sexta-Feira")
	} else if num1 == 7 {
		fmt.Print("Sábado")
	} else if num1 < 1 || num1 > 7 {
		fmt.Print("ERRO")
	}
}
