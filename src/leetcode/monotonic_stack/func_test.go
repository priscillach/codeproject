package monotonic_stack

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func TestTrap(t *testing.T) {
	trapV2([]int{0, 2, 0, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
}
