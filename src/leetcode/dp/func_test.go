package dp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindromeV2(t *testing.T) {
	s := "cbbd"
	longestPalindromeV2("aaaa")
	longestPalindromeV2("aaaa")

	fmt.Println(s[1] == s[2])
}

func TestLongestValidParentheses(t *testing.T) {
	longestValidParentheses(")()())")
}

func TestMaximalSquare(t *testing.T) {
	assert.Equal(t, maximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}), 4)
}

func TestFindLength(t *testing.T) {
	findLength([]int{1, 2, 3, 2, 8}, []int{5, 6, 1, 4, 7})
}
