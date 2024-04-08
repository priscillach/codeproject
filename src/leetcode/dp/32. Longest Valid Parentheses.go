package dp

// o(n^3) exceed time limit
func longestValidParentheses(s string) int {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	maxLen := 0
	for i := 1; i < len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if (i == j+1 || dp[j+1][i-1]) && s[j] == '(' && s[i] == ')' {
				dp[j][i] = true
			}
			for k := j + 2; k <= i-1 && (i-j+1)%2 == 0; k++ {
				if dp[j][k-1] && dp[k][i] {
					dp[j][i] = true
					break
				}
			}
			if dp[j][i] && i-j+1 > maxLen {
				maxLen = i - j + 1
			}
		}
	}
	return maxLen
}
