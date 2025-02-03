package binary_tree

import "leetcode/src/define/mytreenode"

// https://leetcode.com/problems/delete-node-in-a-bst/description/
func deleteNode(root *mytreenode.TreeNode, key int) *mytreenode.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		var prev *mytreenode.TreeNode
		if root.Left != nil {
			cur := root.Left
			if cur.Right == nil {
				cur.Right = root.Right
				return cur
			}
			for cur.Right != nil {
				prev = cur
				cur = cur.Right
			}
			left := root.Left
			right := root.Right
			prev.Right = cur.Left
			cur.Left = left
			cur.Right = right
			return cur
		}

		if root.Right != nil {
			cur := root.Right
			if cur.Left == nil {
				cur.Left = root.Left
				return cur
			}
			for cur.Left != nil {
				prev = cur
				cur = cur.Left
			}
			left := root.Left
			right := root.Right
			prev.Left = cur.Right
			cur.Left = left
			cur.Right = right
			return cur
		}
	}
	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)
	return root
}
