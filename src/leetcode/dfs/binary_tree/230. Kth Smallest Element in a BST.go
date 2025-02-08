package binary_tree

import "leetcode/src/define/mytreenode"

func kthSmallest(root *mytreenode.TreeNode, k int) int {
	var res int
	var kthSmallestDfs func(n *mytreenode.TreeNode)
	kthSmallestDfs = func(n *mytreenode.TreeNode) {
		if n == nil {
			return
		}
		kthSmallestDfs(n.Left)
		k--
		if k == 0 {
			res = n.Val
			return
		}
		kthSmallestDfs(n.Right)
	}
	kthSmallestDfs(root)
	return res
}

func kthSmallestV2(root *mytreenode.TreeNode, k int) int {
	var stack []*mytreenode.TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		k--
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k == 0 {
			return cur.Val
		}
		cur = cur.Right
	}
	return 0
}
