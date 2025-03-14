package main

import (
	"fmt"
	"time"
)

func main() {
	//concorrencia é diferente de paralelismo
	go escrever("Olá mundo")
	go escrever("Programando em go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
