package main

//Exercicio 2
// Uma empresa de meteorologia quer ter um sistema onde possa ter a temperatura, umidade e pressão atmosférica de diferentes lugares.

// Declare 3 variáveis especificando o tipo de dado delas, como valor elas devem possuir a temperatura, umidade e pressão atmosférica de onde você se encontra.
// Imprima os valores no console.
// Quais tipos de dados serão atribuídos a essas variáveis?
// int para temperatura, umidade e pressão atmosférica.

import "fmt"

func main() {
	var temperatura int = 30
	var umidade int = 50
	var pressaoAtmosferica int = 1013
	fmt.Println("Temperatura: ", temperatura)
	fmt.Println("Umidade: ", umidade)
	fmt.Println("Pressão Atmosférica: ", pressaoAtmosferica)
}
