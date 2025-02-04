package binary_tree

import "leetcode/src/define/mytreenode"

// https://leetcode.com/problems/subtree-of-another-tree/description/
func isSubtree(root *mytreenode.TreeNode, subRoot *mytreenode.TreeNode) bool {
	if root == nil {
		return false
	}
	if equal(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func equal(root1 *mytreenode.TreeNode, root2 *mytreenode.TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 != nil && root2 != nil {
		return equal(root1.Left, root2.Left) && equal(root1.Right, root2.Right) && root1.Val == root2.Val
	}
	return false
}
