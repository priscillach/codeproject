package binary_tree

import "leetcode/src/define/mytreenode"

// https://leetcode.com/problems/sum-root-to-leaf-numbers/description/
// finish times: 2
var sum int

func sumNumbers(root *mytreenode.TreeNode) int {
	sum = 0
	sumNumbersCore(root, 0)
	return sum
}

func sumNumbersCore(root *mytreenode.TreeNode, cur int) {
	if root.Left == nil && root.Right == nil {
		sum += cur*10 + root.Val
		return
	}
	if root.Left != nil {
		sumNumbersCore(root.Left, cur*10+root.Val)
	}
	if root.Right != nil {
		sumNumbersCore(root.Right, cur*10+root.Val)
	}
}
