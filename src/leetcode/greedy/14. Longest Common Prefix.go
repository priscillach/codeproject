package greedy

import (
	"leetcode/src/utils/mathhelper"
)

func longestCommonPrefix(strs []string) string {
	first := strs[0]
	for i := 1; i < len(strs); i++ {
		first = longestCommonPrefixCore(first, strs[i])
		if first == "" {
			break
		}
	}
	return first
}

func longestCommonPrefixCore(str1, str2 string) string {
	var res []byte
	for i := 0; i < mathhelper.Min(len(str1), len(str2)); i++ {
		if str1[i] == str2[i] {
			res = append(res, str1[i])
			continue
		}
		break
	}
	return string(res)
}
