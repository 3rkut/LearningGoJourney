// lecture4

package main

import (
	"fmt"
	"unicode/utf8"
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

	message007 := "lets try interesting things."
	fmt.Println(len(message007)) // len function. prints 28.

	message97 := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message97); i++ {
		c4 := message97[i]
		if c4 >= 'a' && c4 <= 'z' {
			fmt.Printf("c before +13:%c \n", c4)
			c4 = c4 + 13
			fmt.Printf("c after +13:%c \n", c4)
			if c4 > 'z' {
				c4 = c4 - 26
				fmt.Printf("c after -26:%c \n", c4)

			}
		}
		fmt.Printf("%c\n", c4)

	}

	question := "¿Cómo estás?" // decoding runes.
	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")
	var1, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes.\n", var1, size)

	for aaa, ccc := range question { // with range usage.
		fmt.Printf("%v %c\n", aaa, ccc)
	}
	fmt.Println("==========")
	for _, ccc := range question {
		fmt.Printf("%c\n", ccc)
	}

	alphabet := "abcdefghijklmnopqrstuvwxyz" // Quick check 9.6.
	for f1, c1 := range alphabet {
		fmt.Printf("\n%v bytes. \n%c", f1, c1)
	}

	quickk := "¿" // other solution
	fmt.Println(len(quickk))

}
