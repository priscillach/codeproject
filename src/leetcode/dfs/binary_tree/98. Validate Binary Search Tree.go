package binary_tree

import (
	"leetcode/src/define/mytreenode"
	"math"
)

func isValidBST(root *mytreenode.TreeNode) bool {
	return isValidBSTCore(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTCore(node *mytreenode.TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return isValidBSTCore(node.Left, min, node.Val) && isValidBSTCore(node.Right, node.Val, max)
}

func isValidBSTV2(root *mytreenode.TreeNode) bool {
	return isValidBSTCoreV2(root, nil, nil)
}

func isValidBSTCoreV2(root *mytreenode.TreeNode, left, right *int) bool {
	if root == nil {
		return true
	}

	if left != nil && root.Val <= *left {
		return false
	}

	if right != nil && root.Val >= *right {
		return false
	}

	return isValidBSTCoreV2(root.Left, left, &root.Val) && isValidBSTCoreV2(root.Right, &root.Val, right)
}
