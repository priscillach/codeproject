package greedy

import "leetcode/src/utils/mathhelper"

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
