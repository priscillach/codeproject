package array_shuffle

import (
	"fmt"
	"testing"
)

func TestArrayShuffle(t *testing.T) {
	solution := Constructor([]int{0, 1, 2, 3})
	fmt.Println(solution.Shuffle())
}
