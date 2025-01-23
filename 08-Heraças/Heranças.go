package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    float64
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	P := pessoa{"Caio Henrique", "Andrade", 21, 1.83}
	fmt.Println(P)

	E := estudante{P, "engenharia da computação", "Universidade veiga de almeida"}

	fmt.Println(E.nome)
}
