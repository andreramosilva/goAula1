// Exercício 2 - Ler arquivo

// A mesma empresa precisa ler o arquivo armazenado, para isso exige que:

// Seja impresso na tela os valores tabelados, com título ( à esquerda para o ID e à direita para o Preço e Quantidade), o preço, a quantidade e abaixo do preço o total deve ser exibido (somando preço por quantidade)

// Exemplo:

// ID                            Preco     Quantidade

// 111223                      30012.00         1

// 444321                    1000000.00         4

// 434321                         50.50         1

//                           4030062.50

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Abrir arquivo
	file, err := os.Open("produtos.csv")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo")
		return
	}
	defer file.Close()

	// Ler arquivo
	scanner := bufio.NewScanner(file)
	var total float64
	fmt.Printf("%-30s%-15s%-15s\n", "ID", "Preço", "Quantidade")
	for scanner.Scan() {
		linha := scanner.Text()
		if strings.HasPrefix(linha, "ID") {
			continue
		}
		valores := strings.Split(linha, ";")
		id, _ := strconv.Atoi(valores[0])
		preco, _ := strconv.ParseFloat(valores[1], 64)
		quantidade, _ := strconv.Atoi(valores[2])
		fmt.Printf("%-30d%-15.2f%-15d\n", id, preco, quantidade)
		total += preco * float64(quantidade)
	}
	fmt.Printf("%-45s%.2f\n", "", total)
}
