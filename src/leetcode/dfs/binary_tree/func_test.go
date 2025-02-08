package binary_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/src/define/mytreenode"
	"leetcode/src/utils/arrayhelper"
	"math"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	root, _ := mytreenode.BuildBinaryTree(arrayhelper.NumsSlice2NumsPtrSlice([]int{3, 5, 1, 6, 2, 0, 8, math.MinInt, math.MinInt, 7, 4}))
	p := root.Left
	q := root.Right

	assert.Equal(t, lowestCommonAncestor(root, p, q).Val, 3)
}

func TestMaxPathSum(t *testing.T) {
	root, _ := mytreenode.BuildBinaryTree(arrayhelper.NumsSlice2NumsPtrSlice([]int{-10, 9, 20, math.MinInt, math.MinInt, 15, 7}))
	assert.Equal(t, maxPathSum(root), 42)
}

func TestIsBalanced(t *testing.T) {
	root := mytreenode.BuildBinaryTreeV2(arrayhelper.NumsSlice2NumsPtrSlice([]int{1, 2, 2, 3, math.MinInt, math.MinInt, 3, 4, math.MinInt, math.MinInt, 4}))
	fmt.Println(isBalanced(root))
}

func TestDiameterOfBinaryTree(t *testing.T) {
	root := mytreenode.BuildBinaryTreeV2(arrayhelper.NumsSlice2NumsPtrSlice([]int{1, 2, 3, 4, 5}))
	diameterOfBinaryTree(root)
}

func TestIsValidBST(t *testing.T) {
	root := mytreenode.BuildBinaryTreeV2(arrayhelper.NumsSlice2NumsPtrSlice([]int{math.MaxInt32}))
	isValidBST(root)
}

func TestHasPathSum(t *testing.T) {
	root := mytreenode.BuildBinaryTreeV2(arrayhelper.NumsSlice2NumsPtrSlice([]int{5, 4, 8, 11, math.MinInt, 13, 4, 7, 2, math.MinInt, math.MinInt, math.MinInt, 1}))
	hasPathSum(root, 22)
}

func TestBuildTree105(t *testing.T) {
	assert.Panics(t, func() {
		BuildTree105([]int{4, -7, -3, -9, 9, 6, 0, -1, 6, -4, -7, -6, 5, -6, 9, -2, -3, -4},
			[]int{-7, 4, 0, -1, 6, -4, 6, 9, -9, 5, -6, -7, -2, 9, -6, -3, -4, -3})
	})
}

func TestTreeToDoublyList(t *testing.T) {
	node := treeToDoublyList(mytreenode.BuildBinaryTreeFromLeetCodeCase("[4,2,5,1,3]"))
	assert.Equal(t, 1, node.Val)
	assert.Equal(t, 5, node.Left.Val)

	node = treeToDoublyList(mytreenode.BuildBinaryTreeFromLeetCodeCase("[7,3,10,2,5,8,11,1,null,4,6,null,9]"))
	assert.Equal(t, 1, node.Val)
	assert.Equal(t, 11, node.Left.Val)
}

func TestMirrorBinaryTree(t *testing.T) {
	root := mytreenode.BuildBinaryTreeFromLeetCodeCase("[4,2,7,1,3,6,9]")
	newRoot := MirrorBinaryTree(root)
	assert.Equal(t, 4, newRoot.Val)
}

func TestFindTargetNode(t *testing.T) {
	root := mytreenode.BuildBinaryTreeFromLeetCodeCase("[7,3,9,1,5]")
	target := 2
	assert.Equal(t, 7, findTargetNode(root, target))
	assert.Equal(t, 7, findTargetNodeV2(root, target))
}

func TestDeleteNode(t *testing.T) {
	deleteNode(mytreenode.BuildBinaryTreeFromLeetCodeCase("[8,0,31,null,6,28,45,1,7,25,30,32,49,null,4,null,null,9,26,29,null,null,42,47,null,2,5,null,12,null,27,null,null,41,43,46,48,null,3,null,null,10,19,null,null,33,null,null,44,null,null,null,null,null,null,null,11,18,20,null,37,null,null,null,null,14,null,null,22,36,38,13,15,21,24,34,null,null,39,null,null,null,16,null,null,23,null,null,35,null,40,null,17]"), 1)
}

func TestFlatten(t *testing.T) {
	flatten(mytreenode.BuildBinaryTreeFromLeetCodeCase("[1,2,5,3,4,null,6]"))
}

func TestKthSmallest(t *testing.T) {
	kthSmallest(mytreenode.BuildBinaryTreeFromLeetCodeCase("[5,3,6,2,4,null,null,1]"), 3)
}
