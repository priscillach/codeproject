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
	nums := []int{0, 2, 1, 5, 4, 2, 3, 1}
	HeapSort(nums)
	fmt.Println(nums)
}

func TestLargestNumber(t *testing.T) {
	fmt.Println(largestNumber([]int{34323, 3432}))
}

func TestReversePairs(t *testing.T) {
	fmt.Println(reversePairs([]int{7, 5, 6, 4}))
	fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
}

func TestFastSort(t *testing.T) {
	nums := []int{0, 2, 1, 5, 4, 2, 3, 1}
	FastSort(nums)
	fmt.Println(nums)
}

func TestRadixSortLSD(t *testing.T) {
	nums := []int{170, 45, 75, 90, 802, 24, 2, 66}
	RadixSortLSD(nums)
	fmt.Println(nums)
}

func TestInsertionSort(t *testing.T) {
	InsertionSort([]int{170, 45, 75, 90, 802, 24, 2, 66})
}

func TestSelectionSort(t *testing.T) {
	nums := []int{170, 45, 75, 90, 802, 24, 2, 66}
	SelectionSort(nums)
	fmt.Println(nums)
}
