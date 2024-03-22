package _map

import "leetcode/src/utils"

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
			p = utils.Max(p, lastIndex[int(s[i])])
		}
		lastIndex[int(s[i])] = i
		maxLen = utils.Max(maxLen, i-p)
	}
	return maxLen
}
