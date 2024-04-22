package dp

import (
	"leetcode/src/utils/mathhelper"
)

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	maxSize := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = mathhelper.Min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
				}
				if dp[i][j]*dp[i][j] > maxSize {
					maxSize = dp[i][j] * dp[i][j]
				}
			}
		}
	}
	return maxSize
}
