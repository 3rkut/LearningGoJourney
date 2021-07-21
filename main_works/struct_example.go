package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	test, err := checkError()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(test)
}

func checkError() (string, error) {
	return "", errors.New("Error detected!")

}
