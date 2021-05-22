package main

import "fmt"

func main() {
	hello := []string{sayHello(""), sayHello("joseph")} // slice.
	hello = append(hello, " 50 shades of grey ")

	for i, data := range hello {
		fmt.Println(i, data)
	}
}

func sayHello(name string) string {
	return "Like a diamond..." + name
}
