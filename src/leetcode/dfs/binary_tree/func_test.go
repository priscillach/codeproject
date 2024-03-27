package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"leetcode/src/define/mytreenode"
	"leetcode/src/utils"
	"math"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	root, _ := mytreenode.BuildBinaryTree(utils.NumsSlice2NumsPtrSlice([]int{3, 5, 1, 6, 2, 0, 8, math.MinInt, math.MinInt, 7, 4}))
	p := root.Left
	q := root.Right

	assert.Equal(t, lowestCommonAncestor(root, p, q).Val, 3)
}

func TestMaxPathSum(t *testing.T) {
	root, _ := mytreenode.BuildBinaryTree(utils.NumsSlice2NumsPtrSlice([]int{-10, 9, 20, math.MinInt, math.MinInt, 15, 7}))
	assert.Equal(t, maxPathSum(root), 42)
}
