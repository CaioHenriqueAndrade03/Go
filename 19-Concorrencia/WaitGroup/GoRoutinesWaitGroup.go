package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//concorrencia é diferente de paralelismo

	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("Olá mundo")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em go")
		waitGroup.Done()
	}()

	go func() {
		escrever("Go routine 3")
		waitGroup.Done()
	}()

	go func() {
		escrever("ProgramGo routine 4")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
