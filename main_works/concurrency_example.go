package main

import "fmt"

func main() {
	go countI()
	go countT()
}

func countI() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func countT() {
	for t := 0; t < 10; t++ {
		fmt.Println(t)
	}
}
