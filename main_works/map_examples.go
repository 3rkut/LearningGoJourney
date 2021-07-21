package main

import "fmt"

func main() {
	var myMap map[string]string // string keys, string values.
	myMap = make(map[string]string)
	myMap2 := map[string]string{"Hello2": "World"}
	myMap["hello"] = "world"
	myMap3 := map[string]string{
		"Name":    "Adam",
		"Surname": "Sandler",
	}
	fmt.Printf("My map1: %v\n", myMap)
	fmt.Printf("My map2: %v\n", myMap2)
	delete(myMap2, "Hello2")
	fmt.Printf("My map2 after deleted: %v\n", myMap2)
	fmt.Printf("My map3: %v\n", myMap3)
	delete(myMap3, "Surname")
	fmt.Printf("My map2 after deleted: %v\n", myMap3)
	fmt.Printf("name of map3: %v\n", myMap3["Name"])
	world := myMap3["Name"]
	fmt.Printf("world: %v\n", world)
	missingVar, ok := myMap3["Missing key"] // ok -> boolean.
	if !ok {
		fmt.Println("The key was missing")
	} else {
		fmt.Println(missingVar)
	}

	myMap4 := []string{"apple", "banana", "watermelon"}
	for i := 0; i < len(myMap4); i++ {
		fmt.Println(myMap4[i])
	}

	for i, v := range myMap4 {
		fmt.Println("Index: ", i)
		fmt.Println("Value", v)
	}
}
