package binarysearch

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
	fmt.Println(lengthOfLISV2([]int{10, 9, 2, 5, 3, 4}))
}

func TestMSqrt(t *testing.T) {
	fmt.Println(mySqrt(8))
}

func TestFindPeakElement(t *testing.T) {
	assert.Equal(t, findPeakElement([]int{1, 2, 3, 1}), 2)
}

func TestSearchMatrix(t *testing.T) {
	searchMatrix([][]int{{1, 1}}, 0)
}
