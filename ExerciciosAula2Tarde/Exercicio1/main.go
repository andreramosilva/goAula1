// Exercício 1 - Registro de estudantes

// Uma universidade precisa cadastrar os alunos e gerar uma funcionalidade para imprimir os detalhes dos dados de cada um deles, conforme o exemplo abaixo:

// Nome: [Nome do aluno]
// Sobrenome: [Sobrenome do aluno]
// RG: [RG do aluno]
// Data de admissão: [Data de admissão do aluno]

// Os valores que estão entre parênteses devem ser substituídos pelos dados fornecidos pelos alunos.

// Para isso é necessário gerar uma estrutura Alunos com as variáveis ​​Nome, Sobrenome, RG, Data e que tenha um método de detalhamento

package main

import (
	"fmt"
)

type Alunos struct {
	Nome      string
	Sobrenome string
	RG        string
	Data      string
}

func (a Alunos) detalhamento() {
	fmt.Println("Nome:", a.Nome)
	fmt.Println("Sobrenome:", a.Sobrenome)
	fmt.Println("RG:", a.RG)
	fmt.Println("Data de admissão:", a.Data)
}

func main() {
	aluno1 := Alunos{
		Nome:      "João",
		Sobrenome: "Silva",
		RG:        "123456",
		Data:      "01/01/2021",
	}

	aluno1.detalhamento()
}
