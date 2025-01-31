package dp

import (
	"leetcode/src/utils/mathhelper"
)

// https://leetcode.com/problems/longest-common-subsequence/description/
// finish time: 2
func longestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = mathhelper.Max(dp[i][j], dp[i-1][j-1]+1)
			}
			dp[i][j] = mathhelper.Max(dp[i][j], dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[len1][len2]
}
