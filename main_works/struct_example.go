package main

import "fmt"

type exampleStruct struct {
	name string
	id   int
}

func main() {
	person1 := exampleStruct{"Lucifer", 1}
	fmt.Println(person1.Hellothere())
}

func (es exampleStruct) Hellothere() string {
	return fmt.Sprintf("welcome to the club, %s\n", es.name)
}
