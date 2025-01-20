package _map

import (
	"leetcode/src/utils/mathhelper"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
// finish times: 2
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
			p = mathhelper.Max(p, lastIndex[int(s[i])])
		}
		lastIndex[int(s[i])] = i
		maxLen = mathhelper.Max(maxLen, i-p)
	}
	return maxLen
}
