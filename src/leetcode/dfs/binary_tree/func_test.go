package binary_tree

import (
	"fmt"
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

func TestIsBalanced(t *testing.T) {
	root := mytreenode.BuildBinaryTreeV2(utils.NumsSlice2NumsPtrSlice([]int{1, 2, 2, 3, math.MinInt, math.MinInt, 3, 4, math.MinInt, math.MinInt, 4}))
	fmt.Println(isBalanced(root))
}

func TestDiameterOfBinaryTree(t *testing.T) {
	root := mytreenode.BuildBinaryTreeV2(utils.NumsSlice2NumsPtrSlice([]int{1, 2, 3, 4, 5}))
	diameterOfBinaryTree(root)
}

func TestIsValidBST(t *testing.T) {
	root := mytreenode.BuildBinaryTreeV2(utils.NumsSlice2NumsPtrSlice([]int{math.MaxInt32}))
	isValidBST(root)
}

func TestHasPathSum(t *testing.T) {
	root := mytreenode.BuildBinaryTreeV2(utils.NumsSlice2NumsPtrSlice([]int{5, 4, 8, 11, math.MinInt, 13, 4, 7, 2, math.MinInt, math.MinInt, math.MinInt, 1}))
	hasPathSum(root, 22)
}
