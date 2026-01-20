// Package main
package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)

	*h = old[0 : n-1]

	return old[n-1]
}

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 2}

	myHeap := &IntHeap{}

	for _, num := range arr {
		heap.Push(myHeap, num)
	}

	heap.Init(myHeap)
	heap.Push(myHeap, 0)

	min := heap.Pop(myHeap)

	fmt.Println("min: ", min)

	heap.Push(myHeap, 3)

	fmt.Println("pop 1: ", heap.Pop(myHeap), "pop 2: ", heap.Pop(myHeap), "pop 3: ", heap.Pop(myHeap))
}
