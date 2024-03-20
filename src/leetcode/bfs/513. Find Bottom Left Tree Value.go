package bfs

import "leetcode/src/define/mytreenode"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findBottomLeftValue(root *mytreenode.TreeNode) int {
	queue := make([]*mytreenode.TreeNode, 0)
	queue = append(queue, root)
	ans := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			if i == 0 {
				ans = curr.Val
			}
		}
	}
	return ans
}
