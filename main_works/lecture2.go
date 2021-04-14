package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var num = rand.Intn(10) // random sayi olusturduk.
	fmt.Println(num)

	var num2 = rand.Intn(20)
	fmt.Println(num2)

	// var command = "asda"

	var oda = "soguk"

	switch {
	case oda == "sicak":
		fmt.Println("oda sicak")
	case oda == "soguk":
		fmt.Println("oda soguk!")
		fallthrough // bu sonraki case'i calistirir.
	case oda == "cok soguk!":
		fmt.Println("ustune bir seyler giy!")
	}

	a := 200
	if a != 201 {
		fmt.Println("hello")
	}

}
