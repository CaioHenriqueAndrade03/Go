package main

import "fmt"

var diaSemana string

func diaDiaSemana(numero int) string {
	switch numero {
	case 1:
		diaSemana = "Domingo"
		fallthrough
	case 2:
		diaSemana = "Segunda-feira"
	case 3:
		diaSemana = "Terça-feira"
	case 4:
		diaSemana = "Quarta-feira"
	case 5:
		diaSemana = "Quinta-feira"
	case 6:
		diaSemana = "Sexta-feira"
	case 7:
		diaSemana = "Sabado"
	default:
		diaSemana = "numero invalido"
	}
	return diaSemana

}

func diaDiaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda"
	case numero == 3:
		return "Terça"
	case numero == 4:
		return "Quarta"
	case numero == 5:
		return "Quinta"
	case numero == 6:
		return "Sexta"
	case numero == 7:
		return "Sabado"
	default:
		return "numero invalido"
	}

}
func main() {
	fmt.Println("Switch")

	dia := diaDiaSemana(1)

	fmt.Println(dia)
}
