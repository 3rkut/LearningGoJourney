package main

import (
	"fmt"
	"strconv"
)

func main() {

	downit := "lets use plus operator " + "2" + " times."
	fmt.Println(downit)
	var (
		bh float64 = 32767
		h          = int16(bh)
		// c = "10" - 1 quickcheck 10.1 , with an error:
		//The range keyword can decode a UTF-8 encoded string into runes
		countdown = 1
		str       = "hmm there is " + strconv.Itoa(countdown) + " string in here."
	)
	fmt.Println(h)
	fmt.Println(str)

}
