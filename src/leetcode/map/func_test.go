package _map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	lengthOfLongestSubstring("sbacdafgs")
}

func TestLongestConsecutive(t *testing.T) {
	assert.Equal(t, 9, longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
