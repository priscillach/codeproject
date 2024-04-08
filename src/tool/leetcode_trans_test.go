package tool

import (
	"fmt"
	"strings"
	"testing"
)

func TestTransLCP(t *testing.T) {
	transCode := `
func longestValidParentheses(s string) int {
	var stack []int
	maxLen := 0
	for i := 0; i < len(s); i++ {
		stack = append(stack, i)
		for len(stack) >= 2 && s[stack[len(stack)-1]] == ')' && s[stack[len(stack)-2]] == '(' {
			stack = stack[:len(stack)-2]
		}
		if len(stack) > 0 {
			maxLen = utils.Max(maxLen, i - stack[len(stack)-1])
		} else {
			maxLen = i + 1
		}
	}
	return maxLen
}
`
	transCode = strings.ReplaceAll(transCode, "mytreenode.", "")
	transCode = strings.ReplaceAll(transCode, "mylinkednode.", "")
	if strings.Contains(transCode, "utils.Max") {
		transCode = strings.ReplaceAll(transCode, "utils.Max", "max")
		transCode += `
func max(nums... int) int {
	max := math.MinInt32
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
`
	}
	if strings.Contains(transCode, "utils.Min") {
		transCode = strings.ReplaceAll(transCode, "utils.Min", "min")
		transCode += `
func min(nums... int) int {
	min := math.MaxInt32
	for i := range nums {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
`
	}
	fmt.Println(transCode)
}
