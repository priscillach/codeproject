package binarysearch

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	a := []int{1, 2, 3, 3, 3, 3, 4, 5, 6}
	fmt.Println(searchTheFirst(a, 3))
	fmt.Println(searchTheLast(a, 3))
}

func TestSearchRotate(t *testing.T) {
	a := []int{3, 1}
	fmt.Println(search(a, 3))
}

func TestLengthOfLIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{10, 9, 11, 13, 2, 2, 3, 4, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLISV2([]int{10, 9, 11, 13, 2, 2, 3, 4, 5, 3, 7, 101, 18}))
}
