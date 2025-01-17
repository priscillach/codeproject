package dp

import (
	"leetcode/src/utils/mathhelper"
)

// https://leetcode.com/problems/house-robber/description/
// finish times: 2
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
		// or
		//newUnpick := max(unpick, pick)
		//pick = unpick + nums[i]
		//unpick = newUnpick
	}
	return mathhelper.Max(unPicked, picked)
}
