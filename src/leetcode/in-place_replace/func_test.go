package in_place_replace

import (
	"fmt"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})
}

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	rotate(nums, 4)
	fmt.Println(nums)
}
