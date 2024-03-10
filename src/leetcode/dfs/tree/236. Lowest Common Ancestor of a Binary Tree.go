package tree

import "leetcode/src/define/treenode"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *treenode.TreeNode) *treenode.TreeNode {
	if root == nil {
		return nil
	}
	left := lowestCommonAncestor(root.Left, p, q)
	if left == p && root == q || left == q && root == p {
		return root
	}
	right := lowestCommonAncestor(root.Right, p, q)
	if right == p && root == q || right == q && root == p {
		return root
	}
	if left == p && right == q || left == q && right == p {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	if root == p || root == q {
		return root
	}
	return nil
}
