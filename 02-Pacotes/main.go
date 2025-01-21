package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo no arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("chrj@gmail.com")
	fmt.Println(erro)
}
