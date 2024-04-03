// Exercício 5 - Cálculo da quantidade de alimento

// Um abrigo de animais precisa descobrir quanta comida comprar para os animais de estimação. No momento eles só têm tarântulas, hamsters, cachorros e gatos, mas a previsão é que haja muito mais animais para abrigar.

// Cães precisam de 10 kg de alimento
// Gatos 5 kg
// Hamster 250 gramas.
// Tarântula 150 gramas.

// Precisamos:

// Implementar uma função Animal que receba como parâmetro um valor do tipo texto com o animal especificado e que retorne uma função com uma mensagem (caso não exista o animal)
// Uma função para cada animal que calcule a quantidade de alimento com base na quantidade necessária do animal digitado.
// Exemplo:
// const (

//    dog = "dog"

//    cat = "cat"

// )

// ...

// animalDog, msg := Animal(dog)

// animalCat, msg := Animal(cat)

// ...

// var amount float64

// amount+= animaldog(5)

// amount+= animalCat(8)

package main

func Animal(animal string) (func(int) float64, string) {
	switch animal {
	case "dog":
		return dog, ""
	case "cat":
		return cat, ""
	default:
		return nil, "Animal não encontrado"
	}
}

func dog(quantity int) float64 {
	return float64(quantity * 10)
}

func cat(quantity int) float64 {
	return float64(quantity * 5)
}

func main() {
	animalDog, msg := Animal("dog")
	animalCat, msg := Animal("cat")

	var amount float64

	amount += animalDog(5)
	amount += animalCat(8)
}
