package greedy

import (
	"fmt"
	"testing"
)

func TestJumpGame(t *testing.T) {
	canJump([]int{1, 2, 3})
}

func TestFindPrimeFactors(t *testing.T) {
	fmt.Println(findPrimeFactors(56))
	fmt.Println(findPrimeFactors(30))
	fmt.Println(findPrimeFactors(17))
}

func TestReverseNumber(t *testing.T) {
	fmt.Println(reverseNumber(-1))
	fmt.Println(reverseNumber(-123000))
	fmt.Println(reverseNumber(1534236469))
}

func TestReverse(t *testing.T) {
	reverse(1534236469)
}
