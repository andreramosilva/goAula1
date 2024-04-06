// Exercício 2 - Imposto sobre o salário #2

// Faça a mesma coisa que no exercício anterior, mas reformule o código para que, em vez de “Error()”, seja implementado  “errors.New()”.

package main

import (
	"errors"
	"fmt"
)

func main() {
	salario := 10000
	if salario < 15000 {
		err := errors.New("erro: O salário digitado não está dentro do valor mínimo")
		fmt.Println(err)
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}
