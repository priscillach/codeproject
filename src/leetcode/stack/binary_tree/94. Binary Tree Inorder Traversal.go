package binary_tree

import "leetcode/src/define/mytreenode"

func inorderTraversal(root *mytreenode.TreeNode) []int {
	cur := root
	var res []int
	var stack []*mytreenode.TreeNode
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
