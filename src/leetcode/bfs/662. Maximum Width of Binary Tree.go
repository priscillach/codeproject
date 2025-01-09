package bfs

import (
	"leetcode/src/define/mytreenode"
)

// https://leetcode.com/problems/maximum-width-of-binary-tree/submissions/1503419371/
// finish times: 2
func widthOfBinaryTree(root *mytreenode.TreeNode) int {
	var queue []*mytreenode.TreeNode
	root.Val = 0
	queue = append(queue, root)
	maxWidth := 1
	for len(queue) > 0 {
		size := len(queue)
		var left, right int
		for i := 0; i < size; i++ {
			pop := queue[0]
			if i == 0 {
				left = pop.Val
			}
			if i == size-1 {
				right = pop.Val
			}
			queue = queue[1:]
			if pop.Left != nil {
				pop.Left.Val = pop.Val*2 + 1
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				pop.Right.Val = pop.Val*2 + 2
				queue = append(queue, pop.Right)
			}
		}
		if maxWidth < right-left+1 {
			maxWidth = right - left + 1
		}
	}
	return maxWidth
}
