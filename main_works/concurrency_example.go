package main

import (
	"fmt"
	"sync"
)

func main() { // func main is first Goroutine.
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		countI()
		wg.Done()
	}()
	wg.Wait()
}

func countI() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
