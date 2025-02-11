package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total

}

func escrever(texto string, numeros ...int) {
	for _, numeros := range numeros {
		fmt.Println(texto, numeros)
	}
}

func main() {
	fmt.Println(`Função variatica`)
	totalDaSoma := soma(0, 10, 20, 30, 40, 50)
	fmt.Println(totalDaSoma)
}
