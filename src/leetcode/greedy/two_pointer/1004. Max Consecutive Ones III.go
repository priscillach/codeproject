package two_pointer

import "leetcode/src/utils/mathhelper"

func longestOnes(nums []int, k int) int {
	maxOnes := 0
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			k--
		}
		right++
		if k == 0 {
			maxOnes = mathhelper.Max(maxOnes, right-left)
		}
		if k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}
	}
	return mathhelper.Max(maxOnes, right-left) // target nums = [0,0,0,1], k = 4
}
