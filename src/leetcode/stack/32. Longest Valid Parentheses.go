package stack

import (
	"leetcode/src/utils/mathhelper"
)

// https://leetcode.com/problems/longest-valid-parentheses/description/
// finish times: 2
func longestValidParentheses(s string) int {
	var stack []int
	maxLen := 0
	for i := 0; i < len(s); i++ {
		stack = append(stack, i)
		for len(stack) >= 2 && s[stack[len(stack)-1]] == ')' && s[stack[len(stack)-2]] == '(' {
			stack = stack[:len(stack)-2]
		}
		if len(stack) > 0 {
			maxLen = mathhelper.Max(maxLen, i-stack[len(stack)-1])
		} else {
			maxLen = i + 1
		}
	}
	return maxLen
}

func longestValidParenthesesV2(s string) int {
	var stack []int
	maxLen := 0
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && s[stack[len(stack)-1]] == '(' && s[i] == ')' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				maxLen = i + 1
			} else if maxLen < i-stack[len(stack)-1] {
				maxLen = i - stack[len(stack)-1]
			}
		} else {
			stack = append(stack, i)
		}
	}
	return maxLen
}
