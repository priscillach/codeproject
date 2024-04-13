package binary_tree

import (
	"leetcode/src/define/mytreenode"
)

func hasPathSum(root *mytreenode.TreeNode, targetSum int) bool {
	return hasPathSumCore(root, targetSum, 0)
}

func hasPathSumCore(root *mytreenode.TreeNode, targetSum, sum int) bool {
	if root == nil {
		return false
	}
	sum += root.Val
	if root.Left == nil && root.Right == nil && sum == targetSum {
		return true
	}
	return hasPathSumCore(root.Left, targetSum, sum) || hasPathSumCore(root.Right, targetSum, sum)
}
