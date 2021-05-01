// chapter1.1.go
// malwarehenri

package main

import (
	"fmt"
)

func main() {

	var p *int
	a := 10
	p = &a

	fmt.Println("a: ", a)
	fmt.Println("&a: ", &a)
	fmt.Println("p: ", p)
	fmt.Println("&p: ", &p)
	fmt.Println("*p: ", *p)

	*p = 30 // change the *p and check a variable.
	fmt.Println("a: ", a)
}
