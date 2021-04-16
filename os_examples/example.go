package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

func main() {

	var (
		speed = 100
		car   = "Lamborghini"
	)
	fmt.Println(speed)
	fmt.Println(car)
	const number = 1
	fmt.Println(number)

	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)

	/*data, err := os.ReadFile("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)*/ // for read a.txt file

	// we write this string in b.txt file with WriteFile
	// reference: https://pkg.go.dev/os?utm_source=gopls#WriteFile
	err := os.WriteFile("b.txt", []byte("This is WriteFile func example!!!"), 0666)
	if err != nil {
		log.Fatal(err)
	}

	// with ReadFile function, we read what we wrote.
	// reference: https://pkg.go.dev/os?utm_source=gopls#ReadFile
	data, err := os.ReadFile("b.txt")
	os.Stdout.Write(data)

}
