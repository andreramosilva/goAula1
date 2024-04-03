package main

import "fmt"

func calcularSalario(horasMensais int64, categoria string) float64 {
	var salario float64 = 0.0
	if categoria == "A" {
		salario = 3000
		if horasMensais > 160 {
			return (float64(horasMensais) * salario) * 0.5
		}

	}
	if categoria == "B" {
		salario = 1500
		if horasMensais > 160 {
			return (float64(horasMensais) * salario) * 0.2
		}
	}
	if categoria == "C" {
		salario = 1000
	}
	return (float64(horasMensais) * salario)
}

func main() {
	fmt.Println(calcularSalario(162, "C"))
	fmt.Println(calcularSalario(176, "B"))
	fmt.Println(calcularSalario(172, "A"))

}
