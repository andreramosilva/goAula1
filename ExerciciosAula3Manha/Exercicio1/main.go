// Exercício 1 - Guardar arquivo
// Uma empresa que vende produtos de limpeza precisa de:

// Implementar uma funcionalidade  para guardar um arquivo de texto, com a informação de produtos comprados, separados por ponto e vírgula (csv).
// Deve possuir o ID do produto, preço e a quantidade.
// Estes valores podem ser hardcodeados ou escritos manualmente em uma variável.

package main

import (
	"fmt"
	"os"
)

func main() {
	// Criar arquivo
	file, err := os.Create("produtos.csv")
	if err != nil {
		fmt.Println("Erro ao criar arquivo")
		return
	}
	defer file.Close()

	// Escrever no arquivo
	_, err = file.WriteString("ID;Preço;Quantidade\n")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo")
		return
	}

	_, err = file.WriteString("1;10.50;3\n")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo")
		return
	}

	_, err = file.WriteString("2;5.00;5\n")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo")
		return
	}

	_, err = file.WriteString("3;2.00;10\n")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo")
		return
	}

	fmt.Println("Arquivo criado com sucesso")
}
