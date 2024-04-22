package dp

import (
	"leetcode/src/define/mytreenode"
	"leetcode/src/utils/mathhelper"
)

func rob337(root *mytreenode.TreeNode) int {
	return mathhelper.Max(robCore337(root))
}

func robCore337(root *mytreenode.TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftPicked, leftUnpicked := robCore337(root.Left)
	rightPicked, rightUnpicked := robCore337(root.Right)
	return leftUnpicked + rightUnpicked + root.Val, mathhelper.Max(leftPicked, leftUnpicked) + mathhelper.Max(rightPicked, rightUnpicked)
}
