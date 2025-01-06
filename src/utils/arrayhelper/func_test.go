package arrayhelper

import (
	"fmt"
	"testing"
)

func TestArrayFill(t *testing.T) {
	b := []int{1, 2, 3}
	FillSlice(b, -1)
	fmt.Println(b)
}

func TestRemoveDuplicatesInPace(t *testing.T) {
	nums := []int{5, 2, 5, 3, 5, 2, 2, 3}
	nums = RemoveDuplicatesInPlace(nums)
	fmt.Println(nums)
	nums = []int{5, 2, 5, 3, 5, 2, 2, 3}
	nums = RemoveDuplicatesInPlaceWithOutLibraries(nums)
	fmt.Println(nums)
}
