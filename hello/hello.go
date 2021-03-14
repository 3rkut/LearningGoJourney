package main

import (
    "fmt"
    "example.com/greetings"
)

func main() {
    message1 := greetings.Hello("Gladys")
    fmt.Println(message1)
}
