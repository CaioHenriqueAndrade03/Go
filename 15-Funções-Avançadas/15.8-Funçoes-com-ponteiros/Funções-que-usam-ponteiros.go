package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	fmt.Println("Funçoes que utilizam ponteiros")
	numero := 20
	numeroInvertido := inverterSinal(numero)

	fmt.Println(numeroInvertido)

	novoNumero := 40

	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
