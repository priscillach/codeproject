package two_pointer

import "leetcode/src/utils/mathhelper"

// https://leetcode.com/problems/container-with-most-water/
// finish times: 2
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	var res int
	for left < right {
		res = mathhelper.Max(res, (right-left)*mathhelper.Min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
