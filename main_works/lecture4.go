// lecture4

package main

import (
	"fmt"
)

func main() {

	// var blank string // blank string == ""
	fmt.Println("this is string. \nand this is new line.")
	fmt.Println(`another example \n and another`) // backtick example.
	// backticks indicate a raw string literal.
	fmt.Println(`
		and some strings :)))))
		and another examplesss`) // raw string.
	// var onestr = `C:\Program Files` // example of path te usage.
	// using normal string for define a path fails.

	// note : A byte is an alias for the uint8 type. A rune is an alias for the int32 type.
	// type byte = uint8 , type rune = int32.
	// example of rune define - var pi rune = 900.
	var pi rune = 900
	fmt.Printf("%v", pi)
	fmt.Printf("%c", pi) // prints raw.

	// not := 'A' // or var not rune = 'A'
	// var yildiz byte = '*'

	message := "jamesbond"
	c := message[3] // indexing into a string.
	fmt.Println(c)
	fmt.Printf("%c\n", c) // same thing.
	// you can't modify a string in Go.

	// book example -QC 9.3-

	message2 := "shalom"

	for i := 0; i < 6; i++ {
		a := message2[i]
		fmt.Printf("%c\n", a)
	}

	caesar := 'd'                       // example of manipulating chars with Caesar cipher.
	caesar = caesar + 3                 // a, b, c, (d), e(1), f(2), g(3).
	fmt.Printf("caesar : %c\n", caesar) // prints > g.

	c1 := 'g' // quick check 9.4 , answer.
	c1 = c1 - 'a' + 'A'
	fmt.Printf("result : %c\n", c1)
}
