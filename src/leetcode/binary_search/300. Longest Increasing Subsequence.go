package binarysearch

import (
	"sort"
)

// https://leetcode.com/problems/longest-increasing-subsequence/description/
// finish times: 2
func lengthOfLIS(nums []int) int {
	sub := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > sub[len(sub)-1] {
			sub = append(sub, num)
		} else {
			// Search uses binary search to find and return the smallest index i
			// in [0, n) at which f(i) is true, assuming that on the range [0, n)
			j := sort.Search(len(sub), func(m int) bool { return sub[m] >= num })
			sub[j] = num
		}
	}
	return len(sub)
}

func lengthOfLISV2(nums []int) int {
	sub := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] > sub[len(sub)-1] {
			sub = append(sub, nums[i])
			continue
		}
		left, right := 0, len(sub)-1
		for left < right {
			mid := left + (right-left)>>1
			if sub[mid] >= nums[i] {
				right = mid
			} else {
				left = mid + 1
			}

		}
		sub[left] = nums[i]
	}
	return len(sub)
}
