package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Caio",
		"sobrenome": "Andrade",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Caio",
			"ultimo":   "Andrade",
		},
		"curso": {
			"nome":   "engenharia",
			"campus": "campos1",
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"])

	delete(usuario2, "nome")
	usuario2["signo"] = map[string]string{
		"nome": "Virgem",
	}

	fmt.Println(usuario2)
}
