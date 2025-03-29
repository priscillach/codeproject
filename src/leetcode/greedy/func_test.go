package greedy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestJump(t *testing.T) {
	fmt.Println(jumpV2([]int{3, 0, 0}))
}

func TestFindTimeRangeByTimePoint(t *testing.T) {
	var s, e string
	s, e = findTimeRangeByTimePoint("2024-09-27 16:40:00", "20:00:00 1", "02:00:00 3")
	assert.Equal(t, "2024-09-30 20:00:00", s)
	assert.Equal(t, "2024-10-02 02:00:00", e)
	s, e = findTimeRangeByTimePoint("2024-09-24 16:40:00", "20:00:00 1", "02:00:00 3")
	assert.Equal(t, "2024-09-23 20:00:00", s)
	assert.Equal(t, "2024-09-25 02:00:00", e)
	s, e = findTimeRangeByTimePoint("2024-09-24 16:40:00", "20:00:00 3", "02:00:00 1")
	assert.Equal(t, "2024-09-25 20:00:00", s)
	assert.Equal(t, "2024-09-30 02:00:00", e)
	s, e = findTimeRangeByTimePoint("2024-09-27 16:40:00", "20:00:00 3", "02:00:00 1")
	assert.Equal(t, "2024-09-25 20:00:00", s)
	assert.Equal(t, "2024-09-30 02:00:00", e)
	s, e = findTimeRangeByTimePoint("2024-09-27 16:40:00", "20:00:00 7", "02:00:00 1")
	assert.Equal(t, "2024-09-29 20:00:00", s)
	assert.Equal(t, "2024-09-30 02:00:00", e)
}

func TestIsPalindrome(t *testing.T) {
	isPalindrome(-121)
}

func TestMaxScorePath(t *testing.T) {
	arr := []int{6, -2, 3, 4, 1, -4, 1, 2, -6, -3, 2}
	startScore := 5
	canReach, maxScore, path := maxScorePath(arr, startScore)
	fmt.Println("是否能到达最后一个索引:", canReach)
	fmt.Println("最大分数:", maxScore)
	fmt.Println("路径:", path)
}

func TestCanCompleteCircuit(t *testing.T) {
	canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1})
}
