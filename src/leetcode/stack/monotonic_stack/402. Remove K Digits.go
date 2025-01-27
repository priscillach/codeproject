package monotonic_stack

import (
	"leetcode/src/utils/stringhelper"
	"strconv"
)

// https://leetcode.com/problems/remove-k-digits/description/
func removeKdigits(num string, k int) string {
	var stack []int
	for i := 0; i < len(num); i++ {
		digit := stringhelper.NumByte2Int(num[i])
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > digit {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	var res string
	for len(stack) > 1 && stack[0] == 0 {
		stack = stack[1:]
	}
	// if k > 0, remove digit from tail
	for k > 0 && len(stack) > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	for i := 0; i < len(stack); i++ {
		res += strconv.Itoa(stack[i])
	}
	if res == "" {
		return "0"
	}
	return res
}
