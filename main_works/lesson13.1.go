package main

import (
	"fmt"
	"math"
)

func myfunc1(a, b float64) float64 { // pisagor :)
	return math.Sqrt(a*a + b*b)
}

func myfunc2(x, y int) (int, int) {
	return x + y, x * y
}

func main() {

	var (
		a     = myfunc1(8, 15)
		b, b2 = myfunc2(4, 5)
	)
	fmt.Printf("result: %d\n", a) // 8 - 15 - (17)
	fmt.Println("sum result:", b, "multiply result: ", b2)
}
