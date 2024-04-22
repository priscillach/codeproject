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
