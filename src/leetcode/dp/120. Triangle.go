package dp

import (
	"leetcode/src/utils/mathhelper"
	"math"
)

func minimumTotal(triangle [][]int) int {
	minSum := math.MaxInt
	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = triangle[0][0]
	for i := 0; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if i > 0 {
				if j == 0 {
					dp[i][j] = dp[i-1][j] + triangle[i][j]
				} else if j == len(triangle[i])-1 {
					dp[i][j] = dp[i-1][len(triangle[i])-2] + triangle[i][j]
				} else {
					dp[i][j] = mathhelper.Min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
				}
			}
			if i == len(triangle)-1 {
				minSum = mathhelper.Min(minSum, dp[i][j])
			}
		}
	}
	return minSum
}
