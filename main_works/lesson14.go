package main

import "fmt"

func main() {
	hello := []string{sayHello(""), sayHello("joseph")} // slice.
	hello = append(hello, " 50 shades of grey ")

	for i, data := range hello {
		fmt.Println(i, data)
	}
	hello2 := heck{"Hello Dear Teacher.", sayHello("XD")}
	// run with this command : "go run lesson14.go lesson14.1go"
	fmt.Println(hello2)
}

func sayHello(name string) string {
	return "Like a diamond..." + name
}
