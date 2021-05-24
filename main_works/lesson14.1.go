package main

import "fmt"

type heck []string

func (h heck) print() {
	for i, data := range h {
		fmt.Println(i, data)
	}
}
