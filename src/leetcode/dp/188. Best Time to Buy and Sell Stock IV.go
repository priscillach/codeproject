package dp

import (
	"leetcode/src/utils/mathhelper"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/description/
// finish times: 2
func maxProfit188(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = []int{-prices[0], 0}
	}
	for i := 1; i < len(prices); i++ {
		for j := 0; j < k; j++ {
			if j == 0 {
				dp[j][0] = mathhelper.Max(dp[j][0], -prices[i])
				dp[j][1] = mathhelper.Max(dp[j][1], dp[j][0]+prices[i])
				continue
			}
			dp[j][0] = mathhelper.Max(dp[j][0], dp[j-1][1]-prices[i])
			dp[j][1] = mathhelper.Max(dp[j][1], dp[j][0]+prices[i])
		}
	}
	return dp[k-1][1]
}
