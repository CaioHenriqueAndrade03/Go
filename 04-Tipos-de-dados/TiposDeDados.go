package main

import (
	"errors"
	"fmt"
)

func main() {
	var idade int8 = 21
	var erro error = errors.New("ta errado aqui hein")

	fmt.Println("Minha idade é,", idade)
	fmt.Println(erro)
}
