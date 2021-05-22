package main

import "fmt"

func main() {
	hello := []string{sayHello(""), sayHello("joseph")} // slice.
	fmt.Println(hello)
}

func sayHello(name string) string {
	return "Like a diamond..." + name
}
