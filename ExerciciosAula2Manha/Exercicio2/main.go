package main

import "fmt"

func media(notas []int) float64 {
	sum_notas := 0
	for i := 0; i < len(notas); i++ {
		if notas[i] < 0 {
			return -1.0
		}
		sum_notas += notas[i]
	}
	return float64(sum_notas) / float64(len(notas))
}

func main() {
	notas := []int{10, 10, 10, 10}
	fmt.Println(media(notas))

}
