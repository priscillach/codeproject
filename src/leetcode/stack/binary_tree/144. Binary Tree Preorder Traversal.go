package binary_tree

import "leetcode/src/define/mytreenode"

// https://leetcode.com/problems/binary-tree-preorder-traversal/description/
// finish times: 2
func preorderTraversal(root *mytreenode.TreeNode) []int {
	var res []int
	var stack []*mytreenode.TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			res = append(res, cur.Val)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = cur.Right
	}
	return res
}
