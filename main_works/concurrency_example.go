package main

import (
	"fmt"
	"time"
)

func main() {
	go countI()
	go countT()
	time.Sleep(time.Second)
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
