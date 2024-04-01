package main

// Exercício 1 - Letras de uma palavra
// A Real Academia Brasileira quer saber quantas letras tem uma palavra e então ter cada uma das letras separadamente para soletrá-la. Para isso terão que:

// Crie uma aplicação que tenha uma variável com a palavra e imprima o número de letras que ela contém.
// Em seguida, imprimi cada uma das letras.

import "fmt"

func main() {
	palavra := "golang"
	fmt.Println("Número de letras: ", len(palavra))
	for i := 0; i < len(palavra); i++ {
		fmt.Println(string(palavra[i]))
	}
}
