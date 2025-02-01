package binary_tree

import "leetcode/src/define/mytreenode"

func findTargetNode(root *mytreenode.TreeNode, cnt int) int {
	var res int
	var dfs func(n *mytreenode.TreeNode)
	dfs = func(n *mytreenode.TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Right)
		cnt--
		if cnt == 0 {
			res = n.Val
		}
		dfs(n.Left)
	}
	dfs(root)
	return res
}

func findTargetNodeV2(root *mytreenode.TreeNode, cnt int) int {
	var stack []*mytreenode.TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		}
		cnt--
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cnt == 0 {
			return cur.Val
		}
		cur = cur.Left
	}
	return 0
}
