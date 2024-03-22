package binarysearch

import (
	"sort"
)

func lengthOfLIS(nums []int) int {
	sub := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > sub[len(sub)-1] {
			sub = append(sub, num)
		} else {
			j := sort.Search(len(sub), func(m int) bool { return sub[m] >= num })
			sub[j] = num
		}
	}
	return len(sub)
}

func lengthOfLISV2(nums []int) int {
	var result []int = make([]int, 0, len(nums)/2)
	for i := 0; i < len(nums); i++ {
		if len(result) == 0 || nums[i] > result[len(result)-1] {
			result = append(result, nums[i])
			continue
		}
		var left, right, mid int = 0, len(result), 0
		for left != right {
			mid = left + (right-left)>>1
			if nums[i] > result[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		result[right] = nums[i]
	}
	return len(result)
}
