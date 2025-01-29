package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	fmt.Sprintln("-------------------------------")
	//i := 0

	//for i < 10 {
	//i++
	//time.Sleep(time.Second)
	//fmt.Println("Incrementando i", i)
	//}
	//fmt.Sprintln("-------------------------------")

	//for j := 0; j < 10; j++ {
	//time.Sleep(time.Second)
	//fmt.Println("Incrementando j", j)
	//}

	nomes := [3]string{"Caio", "João", "Matheus"}

	for indice, nome := range nomes {
		fmt.Println(nome, indice)
	}

	for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Caio",
		"Sobrenome": "Andrade",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		time.Sleep(time.Second)
		fmt.Println("OI")
	}
}
