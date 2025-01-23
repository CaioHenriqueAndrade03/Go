package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	cep    int
	rua    string
	numero uint8
}

func main() {
	fmt.Println("arquivos structs")

	//var u usuario
	//u.nome = "Caio Henrique"
	//u.idade = 21
	//fmt.Println(u)

	//u2 := usuario{"CAIO", 21, endereco}
	//fmt.Println(u2)

	//u3 := usuario{nome: "Caio"}
	//fmt.Println(u3)

	endereco := endereco{22775045, "Mario agostinelli", 150}
	fmt.Println(endereco)

	u4 := usuario{"Caio Henrique", 21, endereco}
	fmt.Println(u4)
}
