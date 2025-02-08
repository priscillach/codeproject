package binary_tree

import "leetcode/src/define/mytreenode"

func flatten(root *mytreenode.TreeNode) {
	flattenDfs(root)
}

func flattenDfs(root *mytreenode.TreeNode) (*mytreenode.TreeNode, *mytreenode.TreeNode) {
	if root == nil {
		return nil, nil
	}
	if root.Left == nil && root.Right == nil {
		return root, root
	}
	leftHead, leftTail := flattenDfs(root.Left)
	rightHead, rightTail := flattenDfs(root.Right)
	root.Left = nil
	if leftHead != nil {
		root.Right = leftHead
		if rightHead != nil {
			leftTail.Right = rightHead
			return root, rightTail
		}
	} else if rightHead != nil {
		root.Right = rightHead
		return root, rightTail
	}

	return root, leftTail
}
