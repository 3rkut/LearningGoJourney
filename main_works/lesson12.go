// lesson 12.
package main

import "fmt"

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
}
