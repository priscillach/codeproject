package greedy

import "math"

// https://leetcode.com/problems/palindrome-number/description/
// the reverse of palindrome string or int is the same as the origin
func isPalindrome(x int) bool {
	reverseX := 0
	copyX := x
	if x < 0 {
		return false
	}
	for copyX > 0 {
		mod := copyX % 10
		if mod > math.MaxInt32-10*reverseX {
			return false
		}
		copyX /= 10
		reverseX = reverseX*10 + mod
	}
	return reverseX == x
}
