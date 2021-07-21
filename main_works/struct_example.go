package main

import "fmt"

type exampleStruct struct {
	name string
	id   int
}

func main() {
	person1 := exampleStruct{"Adam", 1}
	// another way.
	// person1 := exampleStruct{
	// name: "Adam",
	// id:   1,
	// }
	fmt.Printf("person1: %v\n", person1)
	fmt.Printf("person1 name: %v\n", person1.name)
	fmt.Printf("person1 id: %v\n", person1.id)
}
