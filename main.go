package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Olá Mundo")
	// var poteDeSorvete string
	// poteDeSorvete = "Feijao"
	// fmt.Println(poteDeSorvete)
	// horas := 20
	// print(horas)

	// files, err := os.ReadDir(".")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// data, err := os.ReadFile("./README.md")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(files)
	// fmt.Println(data)

	var v int = 19
	var p *int

	p = &v
	fmt.Println(p)
	fmt.Println(*p)

}
