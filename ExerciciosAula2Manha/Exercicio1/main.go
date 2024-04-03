package main

import "fmt"

func impostoSalario(salario float64) float64 {
	if salario <= 50000 {
		return 0.00
	} else if salario > 50000 && salario <= 150000 {
		return salario * 0.17
	} else if salario > 150000 {
		return salario * 0.27
	}
	return 0.00
}

func main() {
	// code inside the main function
	fmt.Println(impostoSalario(10000))
	fmt.Println(impostoSalario(50000))
	fmt.Println(impostoSalario(100000))
}
