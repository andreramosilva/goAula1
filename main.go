package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Olá Mundo")
	// var poteDeSorvete string
	// poteDeSorvete = "Feijao"
	// fmt.Println(poteDeSorvete)
	// horas := 20
	// print(horas)

	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}
	data, err := os.ReadFile("./README.md")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(files)
	fmt.Println(data)

}
