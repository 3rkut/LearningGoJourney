package main

import "fmt"

func main() {

	hello2 := heck{"Hello Dear Teacher.", sayHello("XD")}
	// run with this command : "go run lesson14.go lesson14.1go"
	fmt.Println(hello2)
	hello2.print() // call from the lesson14.1.go
}

func sayHello(name string) string {
	return "Like a diamond..." + name
}
