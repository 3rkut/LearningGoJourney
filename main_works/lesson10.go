package main

import (
	"fmt"
	"strconv" // for coverting int to string.
)

func int_to_string(x int) string { // convert int to string with this.
	return strconv.Itoa(x)
}

func main() {

	downit := "lets use plus operator " + "2" + " times."
	fmt.Println(downit)
	var (
		bh float64 = 32767
		h          = int16(bh)
		// c = "10" - 1 quickcheck 10.1 , with an error:
		//The range keyword can decode a UTF-8 encoded string into runes
		countdown = 1
		str       = "hmm there is " + strconv.Itoa(countdown) + " string in here." // strconv package's Atoi function(ascii to int)
		a         = int_to_string(45)
	)
	fmt.Println(h)
	fmt.Println(str)
	fmt.Println(fmt.Sprintf("There is %v more string in here.", countdown)) // another way to convert a number to a string with use Sprintf
	fmt.Printf("a : %v\n", a)
	fmt.Printf(" type of a : %T", a) // prints type of a : string.
}
