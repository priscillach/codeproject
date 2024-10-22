package bfs

import "leetcode/src/define/mytreenode"

func zigzagLevelOrder(root *mytreenode.TreeNode) [][]int {
	var res [][]int
	cnt := 0
	var queue []*mytreenode.TreeNode
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		var level []int
		size := len(queue)
		for i := 0; i < size; i++ {
			pop := queue[0]
			queue = queue[1:]
			if cnt%2 == 0 {
				level = append(level, pop.Val)
			} else {
				level = append([]int{pop.Val}, level...)
			}
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
		res = append(res, level)
		cnt++
	}
	return res
}
