package binary_tree

import "leetcode/src/define/mytreenode"

func isSymmetric(root *mytreenode.TreeNode) bool {
	return isSymmetricCore(root.Left, root.Right)
}

// https://leetcode.com/problems/symmetric-tree/description/
// finish times: 2
func isSymmetricCore(root1, root2 *mytreenode.TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return isSymmetricCore(root1.Left, root2.Right) && isSymmetricCore(root1.Right, root2.Left)
}
