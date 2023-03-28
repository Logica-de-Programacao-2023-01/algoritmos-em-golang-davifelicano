// Faça um algoritmo que imprima a tabuada de multiplicação de 1 a 10 para um número fornecido pelo usuário.
package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Informe um número: ")
	fmt.Scan(&numero)
	for i := 1; i <= 10; i++ {
		fmt.Println(numero * i)

	}
}
