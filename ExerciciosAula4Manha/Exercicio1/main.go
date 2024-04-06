// Exercício 1 - Imposto sobre o salário #1

// Em sua função “main”, defina uma variável chamada “salario”  e atribua um valor do tipo “int”.
// Crie um erro personalizado com uma struct que implemente “Error()” com a mensagem “erro: O salário digitado não está dentro do valor mínimo" em que seja disparado quando “salário” for menor que  15.000. Caso contrário, imprima no console a mensagem “Necessário pagamento de imposto”.

package main

import (
	"fmt"
)

type erroSalario struct {
	salario int
}

func (e erroSalario) Error() string {
	return fmt.Sprintf("erro: O salário digitado não está dentro do valor mínimo")
}

func main() {
	salario := 10000
	if salario < 15000 {
		err := erroSalario{salario}
		fmt.Println(err)
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}
