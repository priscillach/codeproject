package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeString(t *testing.T) {
	decodeString("2[abc]3[cd]ef")
}

func TestCalculate(t *testing.T) {
	assert.Equal(t, 7, calculate227("3+2*2"))
}

func TestCalculate772(t *testing.T) {
	assert.Equal(t, 7, calculate772("1+2 * 3"))
	assert.Equal(t, 21, calculate772("2 *(5+ 5*2) /3 +(6/2+ 8) "))
	assert.Equal(t, -12, calculate772(" (2 +6* 3+5- (3*14/7+2)*5)+3 "))
	assert.Equal(t, 21, calculate772("3 * ( 3 + 3) + 3"))
	assert.Equal(t, 666, calculate772("   (423+(32 * 4+ (4/3 / (312/ 32)+ 32 + 32 - 11) ) * 31 / 23   "))
}

func TestLongestValidParentheses(t *testing.T) {
	assert.Equal(t, 4, longestValidParentheses(")()())"))
}
