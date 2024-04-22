package string_resolve

import (
	"leetcode/src/utils/mathhelper"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	revisions1 := strings.Split(version1, ".")
	revisions2 := strings.Split(version2, ".")
	maxLen := mathhelper.Max(len(revisions1), len(revisions2))
	for i := 0; i < maxLen; i++ {
		var v1, v2 int
		if i < len(revisions1) {
			v1 = itoi(revisions1[i])
		}
		if i < len(revisions2) {
			v2 = itoi(revisions2[i])
		}
		if v1 > v2 {
			return 1
		}
		if v1 < v2 {
			return -1
		}
	}
	return 0
}

func itoi(s string) int {
	res := 0
	for len(s) > 0 && s[0] == '0' {
		s = s[1:]
	}
	for i := 0; i < len(s); i++ {
		res = 10*res + int(s[i]-'0')
	}
	return res
}
