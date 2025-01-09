package bfs

import (
	"leetcode/src/define/mytreenode"
	"math"
)

// https://leetcode.com/problems/check-completeness-of-a-binary-tree/description/
// finish times: 2
func isCompleteTreeV2(root *mytreenode.TreeNode) bool {
	queue := []*mytreenode.TreeNode{root}
	layer := 0
	for len(queue) > 0 {
		size := len(queue)
		lastLayer := int(math.Pow(2, float64(layer))) != size
		hasGap := false
		for i := 0; i < size; i++ {
			first := queue[0]
			queue = queue[1:]
			if lastLayer && (first.Left != nil || first.Right != nil) {
				return false
			}
			if first.Left != nil {
				if hasGap {
					return false
				}
				queue = append(queue, first.Left)
			} else {
				hasGap = true
			}
			if first.Right != nil {
				if hasGap {
					return false
				}
				queue = append(queue, first.Right)
			} else {
				hasGap = true
			}
		}
		layer++
	}

	return true
}

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
