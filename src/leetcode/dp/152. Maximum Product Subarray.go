package dp

import (
	"leetcode/src/utils/mathhelper"
	"math"
)

func maxProduct(nums []int) int {
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = []int{0, 0}
	}
	res := math.MinInt64
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i][0] = nums[i]
			dp[i][1] = nums[i]
		} else if nums[i] > 0 {
			dp[i][0] = mathhelper.Min(nums[i]*dp[i-1][0], nums[i])
			dp[i][1] = mathhelper.Max(nums[i]*dp[i-1][1], nums[i])
		} else if nums[i] < 0 {
			dp[i][0] = mathhelper.Min(nums[i]*dp[i-1][1], nums[i])
			dp[i][1] = mathhelper.Max(nums[i]*dp[i-1][0], nums[i])
		}
		res = mathhelper.Max(res, dp[i][1])
	}
	return res
}
