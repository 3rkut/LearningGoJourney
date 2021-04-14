// lecture3
package main

import (
	"fmt"
	"math/big" // we use big package and learn how to use for big integers.
)

func main() {

	var ( // short syntax for defining more than 1 variable.
		speed  = big.NewInt(245444)
		second = big.NewInt(45545)

		string_example = new(big.Int) // like methods.
	)
	string_example.SetString("21000000000", 10)

	fmt.Printf("string: %s\n", string_example)
	fmt.Printf("speed: %d\n", speed)
	fmt.Printf("second: %d\n", second)

	const length = 10                          // init length value.
	const mymessage = "mars is a great place." // const variables can not change.

	fmt.Println("My message:", mymessage)

}
