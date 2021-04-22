package main

import "fmt"

func main() {

	type celcius float64

	var temperature celcius = 20
	fmt.Println(temperature)

	const degrees = 20
	temperature = degrees
	temperature += 10

	var warmUp float64 = 10
	temperature += warmUp // not same type.

}
