// chapter1.1.go
// malwarehenri

package main

import (
	"fmt"
)

func example(a *int) {
	*a = 80
}

func main() {

	var p *int
	a := 10
	p = &a

	fmt.Println("a: ", a)
	fmt.Println("&a: ", &a)
	fmt.Println("p: ", p)
	fmt.Println("&p: ", &p)
	fmt.Println("*p: ", *p)

	*p = 30               // change the *p and check a variable.
	fmt.Println("a: ", a) // prints 30.

	example(&a)           // a = 30, we give a's address , because parameter must be pointer.
	fmt.Println("a: ", a) // prints 80.
}
