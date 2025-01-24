package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Arrays e Slices")

	var array [5]int

	array[2] = 2

	//fmt.Println(array)
	array2 := [5]string{"a", "b", "c", "d", "e"}
	slice2 := array2[1:3]

	fmt.Println(array2)

	//array3 := [...]int{1, 2, 3, 4, 5, 6}
	//fmt.Println(array3)

	slice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}

	slice = append(slice, 22)

	//fmt.Println(slice)

	fmt.Println(slice2)
	//	fmt.Println(reflect.TypeOf(slice))
	//	fmt.Println(reflect.TypeOf(array))

	fmt.Println("---------------")
	slice3 := make([]float32, 10, 11)

	fmt.Println(slice3)

	fmt.Println(len(slice3)) // len mostra o tamanho do array

	fmt.Println(cap(slice3)) // cap mostra a capacidade

	slice3 = append(slice3, 2)
	fmt.Println(slice3)      // cap mostra a capacidade
	fmt.Println(cap(slice3)) // cap mostra a capacidade
	slice3 = append(slice3, 2)
	fmt.Println(slice3)      // cap mostra a capacidade
	fmt.Println(cap(slice3)) // cap mostra a capacidade

}
