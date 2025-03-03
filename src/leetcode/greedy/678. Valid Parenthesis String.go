package greedy

import "leetcode/src/utils/mathhelper"

// https://leetcode.com/problems/valid-parenthesis-string/description/
func checkValidString(s string) bool {
	return checkValidStringDfs(s, 0, 0, 0)
}

// timeout
func checkValidStringDfs(s string, cur, left, right int) bool {
	if cur == len(s) {
		return left == right
	}
	if right > left {
		return false
	}

	if s[cur] == '*' {
		flag := checkValidStringDfs(s, cur+1, left+1, right)
		if flag {
			return true
		}
		flag = checkValidStringDfs(s, cur+1, left, right+1)
		if flag {
			return true
		}
		flag = checkValidStringDfs(s, cur+1, left, right)
		if flag {
			return true
		}
	} else if s[cur] == '(' {
		return checkValidStringDfs(s, cur+1, left+1, right)
	}
	return checkValidStringDfs(s, cur+1, left, right+1)
}

func checkValidStringV2(s string) bool {
	var minCnt, maxCnt int
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			minCnt--
			maxCnt++
		} else if s[i] == '(' {
			minCnt++
			maxCnt++
		} else {
			minCnt--
			maxCnt--
		}
		if maxCnt < 0 {
			return false
		}
		minCnt = mathhelper.Max(0, minCnt)
	}
	return minCnt == 0
}
