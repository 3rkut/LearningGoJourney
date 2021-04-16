// lecture1
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %d\n", i)
	}

	var count = 0
	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}
	fmt.Println(count)
	if num := rand.Intn(3); num == 0 { // short decleration in a if statement
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}

	switch xd := "fuasd"; xd { // short switch yapisi yazimi.

	case "deneme":
		fmt.Println("lets try something")

	case "fuck":
		fmt.Println("molly")

	case "holy":
		fmt.Println("holy")
	default:
		fmt.Println("muck this shit!")
	}

	var fiyat float64 // atama yapilmazsa, go'da default deger 0'dir.
	fmt.Println(fiyat)

	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%07.2f\n", third)

	piggyBank := 0.0

	for i := 0; i < 11; i++ {
		piggyBank += 0.1
	}
	fmt.Println(piggyBank)
	/*
		switch num := rand.Intn(10); num { // short decleration in a switch statement
		case 0:
			fmt.Println("Space Adventures")
		case 1:
			fmt.Println("SpaceX")
		case 2:
			fmt.Println("Virgin Galactic")
		default:
			fmt.Println("Random spaceline #", num)
		} */
}
