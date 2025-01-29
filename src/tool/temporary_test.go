package tool

import (
	"math"
	"sort"
	"testing"
)

func threeSumClosest(nums []int, target int) int {
	minDif := math.MaxInt
	sort.Ints(nums)
	var minSum int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return target
			} else if sum < target {
				if minDif > target-sum {
					minSum = sum
					minDif = target - sum
				}
				left++
			} else {
				if minDif > sum-target {
					minSum = sum
					minDif = sum - target
				}
				right--
			}
		}
	}
	return minSum
}

func TestThreeSumClosest(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	target := 1
	threeSumClosest(nums, target)
}
