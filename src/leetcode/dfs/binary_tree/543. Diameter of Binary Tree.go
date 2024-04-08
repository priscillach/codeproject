package binary_tree

import (
	"leetcode/src/define/mytreenode"
	"leetcode/src/utils"
)

var diameterOfBinaryTreeMax int

func diameterOfBinaryTree(root *mytreenode.TreeNode) int {
	diameterOfBinaryTreeMax = 0
	diameterOfBinaryTreeCore(root)
	return diameterOfBinaryTreeMax
}

func diameterOfBinaryTreeCore(root *mytreenode.TreeNode) int {
	if root == nil {
		return 0
	}
	depthLeft := diameterOfBinaryTreeCore(root.Left)
	depthRight := diameterOfBinaryTreeCore(root.Right)
	diameterOfBinaryTreeMax = utils.Max(diameterOfBinaryTreeMax, depthLeft+depthRight)
	return utils.Max(depthLeft, depthRight) + 1
}
