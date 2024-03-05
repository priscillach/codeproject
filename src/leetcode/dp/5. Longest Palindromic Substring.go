package dp

func longestPalindrome(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}
	for j := 0; j < l; j++ {
		dp[j][j] = 1
	}
	maxSubStr := s[:1]
	for i := 1; i <= l-1; i++ {
		for j := 0; j < l-i; j++ {
			if s[j] == s[j+i] && (i == 1 || dp[j+1][j+i-1] == 1) {
				dp[j][j+i] = 1
				if i+1 > len(maxSubStr) {
					maxSubStr = s[j : j+i+1]
				}
			}
		}
	}
	return maxSubStr
}

func longestPalindromeV2(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}
	for j := 0; j < l; j++ {
		dp[j][j] = 1
	}
	maxSubStr := s[:1]
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if s[i] == s[j] && (i-j == 1 || dp[j+1][i-1] == 1) {
				dp[j][i] = 1
				if i-j+1 > len(maxSubStr) {
					maxSubStr = s[j : i+1]
				}
			}
		}
	}
	return maxSubStr
}

// V3 Manacher‘s Algorithm
