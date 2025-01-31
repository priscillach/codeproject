package dp

import (
	"leetcode/src/utils/mathhelper"
)

// https://leetcode.com/problems/coin-change/description/
// finish times: 2
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] == -1 {
				continue
			} else if dp[j] == -1 {
				dp[j] = dp[j-coins[i]] + 1
			} else {
				dp[j] = mathhelper.Min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	return dp[amount]
}
