package greedy

import "leetcode/src/utils/mathhelper"

func canCompleteCircuit(gas []int, cost []int) int {
	if mathhelper.Sum(gas...) < mathhelper.Sum(cost...) {
		return -1
	}
	// if gas sum >= cost sum, there must be a answer
	pos := 0
	acc := 0
	// find the pos where the left gas of subarray > 0 starts at
	// ref to LeetCode 53. Maximum Subarray
	for i := 0; i < len(gas); i++ {
		acc += gas[i] - cost[i]
		if acc < 0 {
			pos = i + 1
			acc = 0
		}
	}
	return pos
}
