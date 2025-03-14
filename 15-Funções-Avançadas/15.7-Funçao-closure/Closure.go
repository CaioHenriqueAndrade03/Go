package main

import "fmt"

func closure() func() {
	texto := "Dentro da funçao closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "dentro da funçao main"
	fmt.Println(texto)
	funcaoNova := closure()

	funcaoNova()
}
