package main

import (
	"fmt"
)

func main() { // func main is first Goroutine.
	c := make(chan int)
	go countI(1, c)
	for msg := range c {
		fmt.Println(msg)
	}
}

func countI(x int, c chan int) {
	for i := 0; i < 10; i++ {
		c <- x
	}
	close(c)
}
