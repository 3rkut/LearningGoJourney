package main

import "fmt"

func main() {

	hello2 := heck{"Hello Dear Teacher.", sayHello("XD")}
	// run with this command : "go run lesson14.go lesson14.1go"
	fmt.Println(hello2)
}

func sayHello(name string) string {
	return "Like a diamond..." + name
}
