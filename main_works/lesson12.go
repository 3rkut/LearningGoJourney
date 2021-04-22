// lesson 12.
package main

import (
	"fmt"
	"math/rand"
)

func topla_carp(a, b int) (int, int) {
	return a + b, a * b

}

func main() {

	toplam, carpim := topla_carp(3, 13)
	fmt.Println(toplam, carpim)

	deneme := []int{4, 6, 8, 9}
	deneme = append(deneme, 12)
	for key, value := range deneme {
		fmt.Println("index: ", key, "value: ", value)
	}

	myslice := []string{"apple", "orange", "watermelon"}
	for _, i := range myslice {
		fmt.Println(i)
	}

	var deneme1 = []int{10, 12, 13}

	for i := 0; i < 3; i++ {
		fmt.Println("index: ", i, "value: ", deneme1[i])
		deneme1 = append(deneme1, rand.Intn(6))
	}
	fmt.Println("bitti")
	for index2, value2 := range deneme1 {
		fmt.Println("index: ", index2, "value: ", value2)
	}
	fmt.Println("guncel hali", deneme1)

}
