package main

import "fmt"

func funcao1() {
	fmt.Println("executando a funcao 1")
}

func funcao2() {
	fmt.Println("executando a funcao 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. resultado sera retornado")

	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}
func main() {

	fmt.Println(alunoAprovado(7, 8))
}
