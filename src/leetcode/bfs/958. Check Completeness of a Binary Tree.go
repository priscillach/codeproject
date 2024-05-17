package bfs

import "leetcode/src/define/mytreenode"

func isCompleteTree(root *mytreenode.TreeNode) bool {
	var queue []*mytreenode.TreeNode
	var isLastSecLayer bool
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		var hasGap bool
		for i := 0; i < size; i++ {
			first := queue[0]
			queue = queue[1:]
			if first.Left == nil && first.Right != nil {
				return false
			} else if first.Left == nil && first.Right == nil {
				isLastSecLayer = true
				hasGap = true
			} else if first.Left != nil && first.Right != nil {
				if hasGap || isLastSecLayer {
					return false
				}
				queue = append(queue, first.Left)
				queue = append(queue, first.Right)
			} else {
				if isLastSecLayer {
					return false
				}
				isLastSecLayer = true
				if hasGap {
					return false
				}
				queue = append(queue, first.Left)
			}

		}
	}
	return true
}
