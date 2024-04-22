package stack

import (
	"leetcode/src/utils/stringhelper"
)

func calculate772(s string) int {
	var stack []int
	var num int
	var sign byte = '+'
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + stringhelper.NumByte2Int(s[i])
		}
		if s[i] == '(' {
			// find the match ')' when cnt == 0
			var cnt int
			for j := i; j < len(s); j++ {
				if s[j] == '(' {
					cnt++
					continue
				}
				if s[j] == ')' {
					cnt--
				}
				if cnt == 0 {
					// recursively calculate the expression between the '(' and ')' i.e. i + 1 ~ j - 1
					num = calculate772(s[i+1 : j])
					// let the i move to j, next loop i++, then get the next after ')'
					i = j
					break
				}
			}
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '/' || s[i] == '*' || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pop*num)
			case '/':
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pop/num)

			}
			sign = s[i]
			num = 0
		}
	}
	var res int
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}
