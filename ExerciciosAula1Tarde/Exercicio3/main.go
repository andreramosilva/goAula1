// Exercício 3 - A que mês corresponde?

// Faça uma aplicação que contenha uma variável com o número do mês.

// Dependendo do número, imprima o mês correspondente em texto.
// Ocorre a você que isso pode ser resolvido de maneiras diferentes? Qual você escolheria e porquê?
// Ex: 7 de julho.

package main

import "fmt"

func main() {
	var mes int = 7
	switch mes {
	case 1:
		fmt.Println("Janeiro")
	case 2:
		fmt.Println("Fevereiro")
	case 3:
		fmt.Println("Março")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Maio")
	case 6:
		fmt.Println("Junho")
	case 7:
		fmt.Println("Julho")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Setembro")
	case 10:
		fmt.Println("Outubro")
	case 11:
		fmt.Println("Novembro")
	case 12:
		fmt.Println("Dezembro")
	default:
		fmt.Println("Mês inválido")

	}

}
