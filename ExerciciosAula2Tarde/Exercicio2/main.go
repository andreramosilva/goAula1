package main

import (
	"fmt"
)

// Criar uma estrutura “loja” que guarde uma lista de produtos.
// Criar uma estrutura “produto” que guarde o tipo de produto, nome e preço
// Criar uma interface “Produto” que possua o método “CalcularCusto”
// Criar uma interface “Ecommerce” que possua os métodos “Total” e “Adicionar”.

// Será necessário uma função “novoProduto” que receba o tipo de  produto, seu nome e preço, e devolva um Produto.
// Será necessário uma função “novaLoja” que retorne um  Ecommerce.
// Interface Produto:
// Deve possuir o método “CalcularCusto”, onde o mesmo deverá calcular o custo adicional segundo o tipo de produto.
// Interface Ecommerce:
// - Deve possuir o método “Total”, onde o mesmo deverá retornar o preço total com base no custo total dos produtos +  o adicional citado anteriormente (caso a categoria tenha)

//  - Deve possuir o método  “Adicionar”, onde o mesmo deve receber um novo produto e adicioná-lo a lista da loja

type Produto interface {
	CalcularCusto() float64
}

type Ecommerce interface {
	Total() float64
	Adicionar(p Produto)
}

type produto struct {
	tipo  string
	nome  string
	preco float64
}

func (p produto) CalcularCusto() float64 {
	return p.preco
}

type loja struct {
	produtos []Produto
}

func (l *loja) Total() float64 {
	var total float64
	for _, p := range l.produtos {
		total += p.CalcularCusto()
	}
	return total
}

func (l *loja) Adicionar(p Produto) {
	l.produtos = append(l.produtos, p)
}

func main() {

	loja := loja{}

	produto1 := produto{
		tipo:  "Eletrônico",
		nome:  "Smartphone",
		preco: 1500.00,
	}

	produto2 := produto{
		tipo:  "Eletrônico",
		nome:  "Notebook",
		preco: 2500.00,
	}

	loja.Adicionar(produto1)
	loja.Adicionar(produto2)

	fmt.Println(loja.Total())
}
