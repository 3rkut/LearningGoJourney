package main

import (
	"fmt"
)

func main() {

	downit := "lets use plus operator " + "2" + " times."
	fmt.Println(downit)
	var (
		a = 10
		b = "20"
	)

	fmt.Println(a + int(b)) // prints 30. basic example of typecasting.

}
