package dp

import (
	"leetcode/src/define/mytreenode"
	"leetcode/src/utils"
)

func rob337(root *mytreenode.TreeNode) int {
	return utils.Max(robCore337(root))
}

func robCore337(root *mytreenode.TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftPicked, leftUnpicked := robCore337(root.Left)
	rightPicked, rightUnpicked := robCore337(root.Right)
	return leftUnpicked + rightUnpicked + root.Val, utils.Max(leftPicked, leftUnpicked) + utils.Max(rightPicked, rightUnpicked)
}
