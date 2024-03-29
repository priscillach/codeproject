package bfs

import "leetcode/src/define/mytreenode"

func rightSideView(root *mytreenode.TreeNode) []int {
	var queue []*mytreenode.TreeNode
	var res []int
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == size-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}
