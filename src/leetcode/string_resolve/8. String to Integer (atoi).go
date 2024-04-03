package string_resolve

import "leetcode/src/utils"

func myAtoi(s string) int {
	s = trim(s)
	sign := getSign(s)
	if sign == -1 {
		s = s[1:]
	}
	if sign == 1 {
		if s[0] == '+' {
			s = s[1:]
		}
	}
	var res int
	for i := 0; i < len(s); i++ {
		if !isDigit(s[i]) {
			break
		}
		if s[i] == '0' && res == 0 {
			continue
		}
		cur := utils.NumByte2Int(s[i])
		if sign == 1 && cur > 1<<31-1-10*res {
			return 1<<31 - 1
		}
		if sign == -1 && cur > 1<<31-10*res {
			return -1 << 31
		}
		res = 10*res + cur
	}
	return sign * res
}

func trim(s string) string {
	for len(s) > 0 {
		flag := true
		if s[0] == ' ' {
			s = s[1:]
			flag = false
		}
		if len(s) > 0 && !isDigit(s[len(s)-1]) {
			s = s[:len(s)-1]
			flag = false
		}
		if flag {
			break
		}
	}
	return s
}

func isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func getSign(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[0] == '-' {
		return -1
	}
	return 1
}
