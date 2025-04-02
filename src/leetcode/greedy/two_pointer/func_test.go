package two_pointer

import (
	"fmt"
	"testing"
)

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

func TestTriangleNumber(t *testing.T) {
	fmt.Println(triangleNumber([]int{48, 66, 61, 46, 94, 75}))
}

func TestLongestOnes(t *testing.T) {
	longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)
	longestOnes([]int{1, 1, 0, 0, 1, 1, 1}, 1)
}
