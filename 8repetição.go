// Faça um algoritmo que leia um número inteiro positivo e mostre todos os seus divisores.
package main

import "fmt"

func main() {
	var x int
	var i int
	fmt.Print("Informe um número:")
	fmt.Scan(&x)
	for i = 1; ; i++ {
		if x%i == 0 {
			fmt.Println(i, "é divisor de  ", x)
		}
	}
}
