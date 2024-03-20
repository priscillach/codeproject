package bfs

import "leetcode/src/define/mytreenode"

func levelOrder(root *mytreenode.TreeNode) [][]int {
	var res [][]int
	var queue []*mytreenode.TreeNode
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		var level []int
		for i := 0; i < size; i++ {
			pop := queue[0]
			queue = queue[1:]
			level = append(level, pop.Val)
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
