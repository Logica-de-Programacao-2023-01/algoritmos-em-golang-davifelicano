package main

import "fmt"

func main() {
	var sal1 float64
	fmt.Print("Informe seu salário:")
	fmt.Scan(&sal1)
	x := sal1 * 1.10
	y := sal1 * 1.05
	if sal1 < 1000 {
		fmt.Print("Salário com aumento de 10%: ", x)
	} else if sal1 > 1000 {
		fmt.Print("Salário com aumento de 5%: ", y)
	}
}
