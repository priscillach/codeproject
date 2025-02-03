package three_pointer

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	nums := []int{1, 0, 1, 2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
