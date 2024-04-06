// Exercício 4 -  Imposto sobre o salário #4

// Vamos fazer com que nosso programa seja um pouco mais complexo e útil.

// Desenvolva as funções necessárias para permitir que a empresa calcule:
// Salário mensal de um funcionário segundo a quantidade de horas trabalhadas.
// A função receberá as horas trabalhadas no mês e o valor da hora como parâmetro.
// Esta função deve retornar mais de um valor (salário calculado e erro).
// No caso de o salário mensal ser igual ou superior a R$ 20.000, 10% devem ser descontados como imposto.
// Se o número de horas mensais inseridas for menor que 80 ou um número negativo, a função deverá retornar um erro. Deve indicar "erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês".

// Desenvolva o código necessário para cumprir as funcionalidades solicitadas, usando “errors.New()”, “fmt.Errorf()” e “errors.Unwrap()”. Não esqueça de realizar as validações dos retornos de erro em sua função “main()”

package main

import (
	"errors"
	"fmt"
)

func salarioMensal(horasTrabalhadas int, valorHora float64) (float64, error) {
	if horasTrabalhadas < 80 || horasTrabalhadas < 0 {
		return 0, errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
	}

	salario := float64(horasTrabalhadas) * valorHora

	if salario >= 20000 {
		imposto := salario * 0.10
		salario -= imposto
	}

	return salario, nil
}

func main() {
	horasTrabalhadas := 100
	valorHora := 200.0

	salario, err := salarioMensal(horasTrabalhadas, valorHora)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Salário mensal: R$ %.2f\n", salario)
}
