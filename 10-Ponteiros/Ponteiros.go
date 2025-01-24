package main

import "fmt"

func main() {

	// passagem por copia
	var variavel1 int = 10
	var variavel2 = variavel1

	variavel1++

	fmt.Println(variavel1, variavel2)
	fmt.Println("=========================")
	// passagem por ponteiro
	var variavel3 int
	var ponteiro *int

	variavel3 = 101
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro)

}
