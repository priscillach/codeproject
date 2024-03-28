package binary_tree

import "leetcode/src/define/mytreenode"

func inorderTraversal(root *mytreenode.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)
	res = append(res, left...)
	res = append(res, root.Val)
	res = append(res, right...)
	return res
}
