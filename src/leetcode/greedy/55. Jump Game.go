package greedy

import "leetcode/src/utils/mathhelper"

// https://leetcode.com/problems/jump-game/description/
func canJump(nums []int) bool {
	curMaxPos := 0
	curPos := 0
	for curPos < len(nums)-1 {
		curMaxPos = mathhelper.Max(curMaxPos, curPos+nums[curPos])
		if curPos < curMaxPos {
			curPos++
		} else {
			break
		}
	}
	return curPos == len(nums)-1
}

func canJumpV2(nums []int) bool {
	maxIdx := 0
	for i := 0; i < len(nums); i++ {
		maxIdx = mathhelper.Max(maxIdx, i+nums[i])
		if maxIdx == i && nums[i] == 0 && i != len(nums)-1 {
			return false
		}
	}

	return maxIdx >= len(nums)-1
}
