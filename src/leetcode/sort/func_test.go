package _sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{5, 2, 1, 5, 4, 2, 3, 1}
	BubbleSort(nums)
	fmt.Println(nums)
}

func TestHeapSort(t *testing.T) {
	nums := []int{5, 2, 1, 5, 4, 2, 3, 1}
	HeapSort(nums)
	fmt.Println(nums)
}

func TestLargestNumber(t *testing.T) {
	fmt.Println(largestNumber([]int{34323, 3432}))
}
