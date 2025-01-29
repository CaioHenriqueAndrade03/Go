package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	var numero int = 15

	if numero > 15 {
		fmt.Println("maior que 15")
	} else if numero < 15 {
		fmt.Println("o numero é menor 15")

	} else {
		fmt.Println("Numero é igual a 15")
	}

	if outronumero := numero; outronumero > 0 {
		fmt.Println("numero é maior que zero")
	}
}
