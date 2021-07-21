package main

import "fmt"

type exampleStruct struct {
	name string
	id   int
	exampleStruct2
}

type exampleStruct2 struct {
	surname string
}

func main() {
	person1 := exampleStruct{"Lucifer", 1, exampleStruct2{"Morningstar"}}
	fmt.Println(person1.Hellothere())
	person1.id = 10
	fmt.Printf("person1 id: %v\n", person1.id)
	person1.UpdateID(20)
	fmt.Printf("person1 id: %v\n", person1.id)
	fmt.Println(person1.exampleStruct2.surname)
}

func (es exampleStruct) Hellothere() string {
	return fmt.Sprintf("welcome to the club, %s\n", es.name)
}

func (es *exampleStruct) UpdateID(id int) {
	es.id = id
}
