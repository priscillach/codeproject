package mytreenode

import (
	"github.com/stretchr/testify/assert"
	"leetcode/src/utils"
	"math"
	"testing"
)

func TestBuildBinaryTree(t *testing.T) {
	root, err := BuildBinaryTree(utils.NumsSlice2NumsPtrSlice([]int{-10, 9, 20, math.MinInt, math.MinInt, 15, 7}))
	assert.NoError(t, err)
	assert.Equal(t, root.Right.Right.Val, 7)

	_, err = BuildBinaryTree(utils.NumsSlice2NumsPtrSlice([]int{-10, 9, math.MinInt, math.MinInt, math.MinInt, 15, 7}))
	assert.Error(t, err)

	root, err = BuildBinaryTree(utils.NumsSlice2NumsPtrSlice([]int{math.MinInt}))
	assert.Nil(t, root)
}
