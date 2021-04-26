package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

func main() {

	fmt.Println(f() == f()) // prints false.
}
