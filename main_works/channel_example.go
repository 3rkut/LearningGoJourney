package main

import "fmt"

func main() {
	c := make(chan string, 3) // 3 shows how many messages we want to receive.
	c <- "I"
	c <- "told"
	c <- "you man."

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
