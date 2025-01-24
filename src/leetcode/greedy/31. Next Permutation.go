package greedy

import (
	"leetcode/src/leetcode/sort/algo"
)

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		// 1. if there is num[i] < nums[i + 1]
		if nums[i] > nums[i-1] {
			j := findLarger(nums, i, nums[i-1])
			nums[i-1], nums[j] = nums[j], nums[i-1]
			algo.FastSortByRange(nums, i, len(nums)-1)
			return
		}
	}
	// 2. if there is NO num[i] < nums[i + 1]
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

func findLarger(nums []int, left, target int) int {
	minLargerIdx := left
	for i := left + 1; i < len(nums); i++ {
		if nums[i] > target && nums[i] < nums[minLargerIdx] {
			minLargerIdx = i
		}
	}
	return minLargerIdx
}
