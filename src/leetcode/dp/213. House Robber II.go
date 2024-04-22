package dp

import (
	"leetcode/src/utils/mathhelper"
)

func rob213(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return mathhelper.Max(rob198(nums[1:]), rob198(nums[:len(nums)-1]))
}
