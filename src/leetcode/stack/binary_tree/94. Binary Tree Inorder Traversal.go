package binary_tree

import "leetcode/src/define/mytreenode"

// https://leetcode.com/problems/binary-tree-inorder-traversal/description/
// finish times: 2
func inorderTraversal(root *mytreenode.TreeNode) []int {
	var res []int
	var stack []*mytreenode.TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}
