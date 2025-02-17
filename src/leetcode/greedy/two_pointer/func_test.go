package two_pointer

import "testing"

func TestMinWindow(t *testing.T) {
	minWindow("ADOBECODEBANC", "ABC")
}

func TestMinSubArrayLen(t *testing.T) {
	minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
}

func TestIsPalindrome(t *testing.T) {
	isPalindrome("A man, a plan, a canal: Panama")
}

func TestFindDuplicate(t *testing.T) {
	findDuplicateV2([]int{1, 3, 4, 2, 2})
}
