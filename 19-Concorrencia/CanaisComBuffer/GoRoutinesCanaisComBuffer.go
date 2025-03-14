package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Olá mundo"
	canal <- "Olá"

	mensagem := <-canal
	fmt.Println(mensagem)
}
