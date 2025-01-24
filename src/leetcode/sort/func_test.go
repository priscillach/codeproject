package _sort

import (
	"fmt"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	fmt.Println(largestNumber([]int{34323, 3432}))
}

func TestReversePairs(t *testing.T) {
	fmt.Println(reversePairs([]int{7, 5, 6, 4}))
	fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
}
