package greedy

import "math"

// https://leetcode.com/problems/reverse-integer/
func reverse(n int) int {
	if n == 0 {
		return 0
	}
	res := 0
	var sign = 1
	if n < 0 {
		sign = -1
		n = -n
	}
	for n > 0 {
		// use MaxInt32 not MaxInt
		if sign > 0 && res > math.MaxInt32/10 {
			return 0
		}
		if sign < 0 && res > -math.MinInt32/10 {
			return 0
		}
		res = 10*res + n%10
		n /= 10
	}

	return res * sign
}
