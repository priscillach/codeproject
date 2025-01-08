package binary_tree

import "leetcode/src/define/mytreenode"

func MirrorBinaryTree(root *mytreenode.TreeNode) *mytreenode.TreeNode {
	if root == nil {
		return nil
	}
	newNode := &mytreenode.TreeNode{
		Val:   root.Val,
		Left:  MirrorBinaryTree(root.Left),
		Right: MirrorBinaryTree(root.Right),
	}
	return newNode
}
