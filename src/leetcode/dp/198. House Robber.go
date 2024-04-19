package dp

import "leetcode/src/utils"

func rob198(nums []int) int {
	picked, unPicked := 0, 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			picked = nums[i]
			continue
		}
		tmp := unPicked + nums[i]
		unPicked = utils.Max(picked, unPicked)
		picked = tmp
	}
	return utils.Max(unPicked, picked)
}
