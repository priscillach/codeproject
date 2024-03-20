package myheap

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Unsorted array:", arr)
	HeapSortAsc(arr)
	fmt.Println("Sorted asc array:", arr)
	HeapSortDesc(arr)
	fmt.Println("Sorted desc array:", arr)
}
