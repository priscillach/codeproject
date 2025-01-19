package prefix_sum

import (
	"leetcode/src/leetcode/greedy/two_pointer"
	"testing"
)

func TestSubarraySum(t *testing.T) {
	subarraySumV2([]int{1, 2, 3}, 3)
}

func TestNumSubarrayProductLessThanK(t *testing.T) {
	two_pointer.numSubarrayProductLessThanK([]int{100, 100, 2, 6}, 100)
}
