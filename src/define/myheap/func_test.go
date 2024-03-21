package myheap

import (
	"container/heap"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{12, 4, 5, 6, 7, 11, 13, 5, 6, 7}
	fmt.Println("Unsorted array:", arr)
	HeapSortAsc(arr)
	fmt.Println("Sorted asc array:", arr)
	HeapSortDesc(arr)
	fmt.Println("Sorted desc array:", arr)
}

func TestMyHeap(t *testing.T) {
	h := &MinHeap{}
	heap.Init(h)
	arr := []int{12, 4, 5, 6, 7, 11, 13, 5, 6, 7}
	for _, x := range arr {
		heap.Push(h, x)
	}
	assert.Equal(t, heap.Pop(h).(int), 5)
}
