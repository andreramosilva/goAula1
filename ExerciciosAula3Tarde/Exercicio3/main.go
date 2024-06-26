// Exercício 3 - Calcular Preço (Part II)

// Uma empresa nacional é responsável pela venda de produtos, serviços e manutenção.

// Para isso, eles precisam realizar um programa que seja responsável por calcular o preço total dos Produtos, Serviços e Manutenção. Devido à forte demanda e para otimizar a velocidade, eles exigem que o cálculo da soma seja realizado em paralelo por meio de 3 go routines.

// Precisamos de 3 estruturas:

// Produtos: nome, preço, quantidade.
// Serviços: nome, preço, minutos trabalhados.
// Manutenção: nome, preço.

// Precisamos de 3 funções:

// Somar Produtos: recebe um array de produto e devolve o preço total (preço * quantidade).
// Somar Serviços: recebe uma array de serviço e devolve o preço total (preço * média hora trabalhada, se ele não vier trabalhar por 30 minutos, ele será cobrado como se tivesse trabalhado meia hora).
// Somar Manutenção: recebe um array de manutenção e devolve o preço total.

// Os 3 devem ser executados concomitantemente e ao final o valor final deve ser mostrado na tela (somando o total dos 3).

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Produtos
	produtos := []Produto{
		{"Arroz", 10.50, 3},
		{"Feijão", 5.00, 5},
		{"Macarrão", 2.00, 10},
	}

	// Serviços
	servicos := []Servico{
		{"Limpeza", 20.00, 30},
		{"Pintura", 50.00, 60},
		{"Reforma", 100.00, 120},
	}

	// Manutenção
	manutencao := []Manutencao{
		{"Encanamento", 50.00},
		{"Eletricista", 100.00},
	}

	var wg sync.WaitGroup
	wg.Add(3)

	var total float64
	go func() {
		defer wg.Done()
		total += SomarProdutos(produtos)
	}()
	go func() {
		defer wg.Done()
		total += SomarServicos(servicos)
	}()
	go func() {
		defer wg.Done()
		total +=
			SomarManutencao(manutencao)
	}()

	wg.Wait()
	fmt.Println("Total:", total)
}

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

type Servico struct {
	Nome               string
	Preco              float64
	MinutosTrabalhados int
}

type Manutencao struct {
	Nome  string
	Preco float64
}

func SomarProdutos(produtos []Produto) float64 {
	var total float64
	for _, p := range produtos {
		total += p.Preco * float64(p.Quantidade)
	}
	return total
}

func SomarServicos(servicos []Servico) float64 {
	var total float64
	for _, s := range servicos {
		total += s.Preco * float64(s.MinutosTrabalhados) / 30
	}
	return total
}

func SomarManutencao(manutencao []Manutencao) float64 {
	var total float64
	for _, m := range manutencao {
		total += m.Preco
	}
	return total
}
