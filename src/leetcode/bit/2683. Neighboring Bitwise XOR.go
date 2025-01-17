package bit

// https://leetcode.com/problems/neighboring-bitwise-xor/description
func doesValidArrayExist(derived []int) bool {
	var cnt int
	for _, d := range derived {
		if d == 1 {
			cnt++
		}
	}
	return cnt%2 == 0
}
