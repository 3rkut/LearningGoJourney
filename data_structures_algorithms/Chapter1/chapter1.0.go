// chapter1.0.go
// malwarehenri

package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

type IntegerHeap []int

func (iheap IntegerHeap) Len() int { return len(iheap) }

func (iheap IntegerHeap) Less(i, j int) bool { return iheap[i] < iheap[j] }

func (iheap IntegerHeap) Swap(i, j int) { iheap[i], iheap[j] = iheap[j], iheap[i] }

func (iheap *IntegerHeap) Push(heaprintf interface{}) {
	*iheap = append(*iheap, heaprintf.(int))
}

func powerSeries(x int) (int, int) {
	return x * x, x * x * x
}

func (iheap *IntegerHeap) Pop() interface{} {
	var n, x1 int
	var previous IntegerHeap = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

func powerSeries2(x int) (square int, cube int) {
	square = x * x
	cube = square * x
	return
}
func main() {

	var (
		intList        list.List
		square, cube   = powerSeries(4)
		square2, cube2 = powerSeries2(5)
	)
	intList.PushBack(21)
	intList.PushBack(53)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}

	fmt.Println("square: ", square, "cube: ", cube)
	fmt.Println("squar2: ", square2, "cube2: ", cube2)

	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
