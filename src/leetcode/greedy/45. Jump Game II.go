package greedy

import (
	"leetcode/src/utils/mathhelper"
	"math"
)

// https://leetcode.com/problems/jump-game-ii/description/
func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= mathhelper.Min(i+nums[i], len(nums)-1); j++ {
			dp[j] = mathhelper.Min(dp[j], dp[i]+1)
		}
	}
	return dp[len(nums)-1]
}

func jumpV2(nums []int) int {
	cnt, next, cur := 0, 0, 0
	for next < len(nums)-1 {
		farthest := 0
		for i := cur; i <= next; i++ {
			farthest = mathhelper.Max(farthest, nums[i]+i)
		}
		cur = next + 1
		next = farthest
		cnt++
	}
	return cnt
}
