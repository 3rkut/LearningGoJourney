package main

import (
	"fmt"
	"math"
)

func myfunc1(a, b float64) float64 { // pisagor :)
	return math.Sqrt(a*a + b*b)
}
func main() {

	a := myfunc1(8, 15)
	fmt.Printf("result: %d", a)

}
