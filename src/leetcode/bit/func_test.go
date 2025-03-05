package bit

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func TestModWithBitCalculator(t *testing.T) {
	fmt.Println(modWithBitCalculator(10, 4))
	fmt.Println(modWithBitCalculator(10, 3))
	fmt.Println(modWithBitCalculator(9, 3))
	fmt.Println(modWithBitCalculator(10, 0))
	fmt.Println(modWithBitCalculator(10, 10))
}
