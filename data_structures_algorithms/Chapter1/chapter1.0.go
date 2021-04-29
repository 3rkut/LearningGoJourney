// chapter1.0.go
// malwarehenri

package main

import (
	"container/list"
	"fmt"
)

func powerSeries(x int) (int, int) {
	return x * x, x * x * x
}

func powerSeries2(x int) (square int, cube int) {
	square = x * x
	cube = square * x
	return
}
func main() {

	var (
		intList        list.List
		square, cube   = powerSeries(4)
		square2, cube2 = powerSeries2(5)
	)
	intList.PushBack(21)
	intList.PushBack(53)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}

	fmt.Println("square: ", square, "cube: ", cube)
	fmt.Println("squar2: ", square2, "cube2: ", cube2)

}
