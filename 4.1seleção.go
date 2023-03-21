package main

import "fmt"

func main() {
	var altura float64
	var peso float64
	var sexo string

	fmt.Print("Informe sua altura:")
	fmt.Scan(&altura)
	fmt.Print("Informe seu pseo:")
	fmt.Scan(&peso)
	fmt.Print("Informe seu sexo:")
	fmt.Scan(&sexo)
	imc := peso / (altura * altura)
	if imc < 18 {
		fmt.Print("Abaixo do peso", sexo, " IMC=", imc)
	} else if imc > 18 && imc < 25 {
		fmt.Print("Dentro do Peso ideal", sexo, "IMC=,", imc)
	} else {
		fmt.Print("Acima do peso ideal", sexo, "IMC=", imc)
	}
}
