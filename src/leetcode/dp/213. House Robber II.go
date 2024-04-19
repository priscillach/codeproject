package dp

import "leetcode/src/utils"

func rob213(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return utils.Max(robCore213(nums[1:]), robCore213(nums[:len(nums)-1]))
}

func robCore213(nums []int) int {
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
