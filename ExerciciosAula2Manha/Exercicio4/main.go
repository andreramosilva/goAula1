// Exercício 4 - Cálculo de estatísticas

// Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de notas dos alunos de um curso, sendo necessário calcular os valores mínimo, máximo e médio de suas notas.

// Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo, máximo ou média) e que retorna outra função (e uma mensagem caso o cálculo não esteja definido) que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi indicado na função anterior

// Exemplo:

// const (

//    minimum = "minimum"

//    average = "average"

//    maximum = "maximum"

// )

// ...

// minhaFunc, err := operation(minimum)

// averageFunc, err := operation(average)

// maxFunc, err := operation(maximum)

// ...

// minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)

// averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)

// maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

package main

func minimum(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}

	return min
}

func average(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum / len(numbers)
}

func maximum(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	return max
}

func operation(operation string) func(...int) int {
	return func(numbers ...int) int {
		switch operation {
		case "minimum":
			return minimum(numbers...)
		case "average":
			return average(numbers...)
		case "maximum":
			return maximum(numbers...)
		default:
			return 0
		}
	}
}

func main() {
	operation("minimum")(2, 3, 3, 4, 10, 2, 4, 5)
	operation("average")(2, 3, 3, 4, 1, 2, 4, 5)
	operation("maximum")(2, 3, 3, 4, 1, 2, 4, 5)
}
