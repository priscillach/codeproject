package greedy

func reverseNumber(n int) int {
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
		res *= 10
		mod := n % 10
		res += mod
		n /= 10
	}

	return res * sign
}
