package list

// https://leetcode.com/problems/word-break/description/
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]struct{})
	for _, word := range wordDict {
		dict[word] = struct{}{}
	}
	cantSegmented := make([]bool, len(s))
	return wordBreakCore(s, 0, cantSegmented, dict)
}

func wordBreakCore(s string, i int, cantSegmented []bool, dict map[string]struct{}) bool {
	if i == len(s) {
		return true
	}
	if cantSegmented[i] {
		return false
	}
	for j := i + 1; j <= len(s); j++ {
		if _, ok := dict[s[i:j]]; ok {
			if wordBreakCore(s, j, cantSegmented, dict) {
				return true
			}
		}
	}
	cantSegmented[i] = true
	return false
}
