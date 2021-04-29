// chapter1.0.go
// malwarehenri

package main

import (
	"container/list"
	"fmt"
)

func main() {

	var intList list.List
	intList.PushBack(21)
	intList.PushBack(53)

	for element:= intList.Front(); element!= nil; element = element.Next()
	{
		fmt.Println(element.Value.(int))
	}
}




