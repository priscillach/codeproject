package dp

import (
	"leetcode/src/utils/mathhelper"
)

func rob198(nums []int) int {
	picked, unPicked := 0, 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			picked = nums[i]
			continue
		}
		tmp := unPicked + nums[i]
		unPicked = mathhelper.Max(picked, unPicked)
		picked = tmp
	}
	return mathhelper.Max(unPicked, picked)
}
