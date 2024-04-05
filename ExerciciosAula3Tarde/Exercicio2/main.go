// Exercício 2 - E-commerce (Parte II)

// Uma grande empresa de vendas na web precisa adicionar funcionalidades para adicionar produtos aos usuários. Para fazer isso, eles exigem que usuários e produtos tenham o mesmo endereço de memória no main do programa e nas funções.

// Estruturas necessárias:

// Usuário: Nome, Sobrenome, E-mail, Produtos (array de produtos).
// Produto: Nome, preço, quantidade.
// Algumas funções necessárias:

// Novo produto: recebe nome e preço, e retorna um produto.
// Adicionar produto: recebe usuário, produto e quantidade, não retorna nada, adiciona o produto ao usuário.
// Deletar produtos: recebe um usuário, apaga  os produtos do usuário.

package main

import "fmt"

func main() {
	// Criar usuário
	usuario := Usuario{
		Nome:      "João",
		Sobrenome: "Silva",
		Email:     "aaa",
	}

	// Criar produto
	produto := NovoProduto("Arroz", 10.50)

	// Adicionar produto
	AdicionarProduto(&usuario, produto, 3)
	fmt.Println(usuario)

	// Deletar produtos
	DeletarProdutos(&usuario)
	fmt.Println(usuario)
}

type Usuario struct {
	Nome      string
	Sobrenome string
	Email     string
	Produtos  []Produto
}

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func NovoProduto(nome string, preco float64) Produto {
	return Produto{
		Nome:  nome,
		Preco: preco,
	}
}

func AdicionarProduto(u *Usuario, p Produto, quantidade int) {
	p.Quantidade = quantidade
	u.Produtos = append(u.Produtos, p)
}

func DeletarProdutos(u *Usuario) {
	u.Produtos = nil
}
