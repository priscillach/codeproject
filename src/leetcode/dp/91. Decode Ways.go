package dp

import "strconv"

var sumDecodings int

// https://leetcode.com/problems/decode-ways/description/
// dfs timeout
func numDecodings(s string) int {
	sumDecodings = 0
	numDecodingsDfs(s, 0)
	return sumDecodings
}

func numDecodingsDfs(s string, cur int) {
	if cur > len(s) {
		return
	}
	if cur == len(s) {
		sumDecodings++
		return
	}
	if checkValid(s[cur : cur+1]) {
		numDecodingsDfs(s, cur+1)
	}
	if cur+2 <= len(s) && checkValid(s[cur:cur+2]) {
		numDecodingsDfs(s, cur+2)
	}
}

func checkValid(s string) bool {
	if len(s) == 0 || len(s) > 2 {
		return false
	}
	if s == "0" || len(s) == 2 && s[0] == '0' {
		return false
	}
	num, _ := strconv.Atoi(s)
	if num < 1 || num > 26 {
		return false
	}
	return true
}

func numDecodingsDfsV2(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] > '0' && s[i-1] <= '9' {
			dp[i] += dp[i-1]
		}
		if i > 1 && (s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
