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

// https://leetcode.com/problems/longest-palindromic-substring/description/
// finish times: 2
func longestPalindromeV2(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	maxS := s[:1]
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if (i <= j+2 || dp[j+1][i-1]) && s[i] == s[j] {
				dp[j][i] = true
				if i-j+1 > len(maxS) {
					maxS = s[j : i+1]
				}
			}
		}
	}
	return maxS
}

// V3 Manacher‘s Algorithm
