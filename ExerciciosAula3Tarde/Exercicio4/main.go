// Exercício 4 - Ordenamento

// Uma empresa de sistemas precisa analisar que algoritmos de ordenamento utilizar para seus serviços.

// Para eles é necessário instanciar 3 arrays com valores aleatórios não ordenados

// uma matriz de inteiros com 100 valores
// uma matriz de inteiros com 1000 valores
// uma matriz de inteiros com 10.000 valores

// Para instanciar as variáveis, utilize o rand:

// package main

// import (

//    "math/rand"

// )

// func main() {

//    variavel := rand.Perm(100)

//    variave2 := rand.Perm(1000)

//    variave3 := rand.Perm(10000)

// }

// Cada um deve ser ordenado por:

// Inserção
// grupo
// seleção

// Uma rotina para cada execução de classificação

// Tenho que esperar terminar a ordenação de 100 números para continuar com a ordenação de 1000 e depois a ordenação de 10000.

// Por fim, devo medir o tempo de cada um e mostrar o resultado na tela, para saber qual ordenação ficou melhor para cada arranjo.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	variavel := rand.Perm(100)
	variavel2 := rand.Perm(1000)
	variavel3 := rand.Perm(10000)

	start := time.Now()
	go insercao(variavel)
	go insercao(variavel2)
	go insercao(variavel3)
	fmt.Println("Insercao: ", time.Since(start))

	start = time.Now()
	go grupo(variavel)
	go grupo(variavel2)
	go grupo(variavel3)
	fmt.Println("Grupo: ", time.Since(start))

	start = time.Now()
	go selecao(variavel)
	go selecao(variavel2)
	go selecao(variavel3)
	fmt.Println("Selecao: ", time.Since(start))
}

func insercao(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}

func grupo(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func selecao(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
