package two_pointer

import "sort"

// https://leetcode.com/problems/3sum/
// finish times: 2
func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			target := nums[i] + nums[right] + nums[left]
			if target == 0 {
				res = append(res, []int{nums[i], nums[right], nums[left]})
				left, right = left+1, right-1
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if target < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
