package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	var soma = n1 + n2
	return soma
}

func calculosMatematicos(n1, n2 int16) (int16, int16, int16, int16) {

	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2

	return soma, subtracao, multiplicacao, divisao

}

func main() {
	fmt.Println(somar(4, 2))

	var f = func() {
		fmt.Println("funcao f")
	}
	f()

	soma, subtracao, multiplicacao, divisao := calculosMatematicos(10, 15)

	fmt.Println(soma, subtracao, multiplicacao, divisao)
}
