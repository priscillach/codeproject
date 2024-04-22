package stack

import (
	"leetcode/src/utils/stringhelper"
)

func calculate(s string) int {
	var stack []int
	var num int
	var sign byte = '+'
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + stringhelper.NumByte2Int(s[i])
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
