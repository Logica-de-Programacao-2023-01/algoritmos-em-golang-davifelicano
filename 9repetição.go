//	 Faça um algoritmo que leia vários números inteiros e mostre a média aritmética entre eles.
//		A leitura deve ser interrompida quando for lido o valor zero.
package main

import "fmt"

func main() {
	var i int
	var soma int
	cont := 0
	for {
		fmt.Print("Informe um número ,digite 0 para parar a contagem ")
		fmt.Scan(&i)
		if i == 0 {
			break
		}
		soma += i
		cont++
	}
	fmt.Print("A média é : ", soma/cont)
}
