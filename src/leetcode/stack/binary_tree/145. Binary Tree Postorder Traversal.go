package binary_tree

import "leetcode/src/define/mytreenode"

// https://leetcode.com/problems/binary-tree-postorder-traversal/description/
// finish times: 2
func postorderTraversal(root *mytreenode.TreeNode) []int {
	var res []int
	var stack []*mytreenode.TreeNode
	cur := root
	var prev *mytreenode.TreeNode
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		// 没有右子树和刚遍历完右子树，结束可以把父节点加入结果
		if cur.Right == nil || cur.Right == prev {
			res = append(res, cur.Val)
			stack = stack[:len(stack)-1]
			prev = cur
			cur = nil
		} else {
			cur = cur.Right
		}
	}
	return res
}
