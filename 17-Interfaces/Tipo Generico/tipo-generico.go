package main

import "fmt"

func generica(inter interface{}) {
	fmt.Println(inter)
}

func main() {
	generica("oi")
	generica(12)
	generica(true)

	mapa := map[interface{}]interface{}{
		1:            "iue",
		float32(100): true,
		"bom dia":    "boa noite",
		true:         float64(12),
	}

	fmt.Println(mapa)

}
