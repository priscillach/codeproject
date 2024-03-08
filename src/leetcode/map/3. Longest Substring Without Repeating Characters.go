package _map

const LEN = 256

func lengthOfLongestSubstring(s string) int {
	lastIndex := [LEN]int{}
	for i := range lastIndex {
		lastIndex[i] = -1
	}
	p := -1
	maxLen := 0
	for i := 0; i < len(s); i++ {
		if lastIndex[int(s[i])] != -1 {
			p = max(p, lastIndex[int(s[i])])
		}
		lastIndex[int(s[i])] = i
		maxLen = max(maxLen, i-p)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
