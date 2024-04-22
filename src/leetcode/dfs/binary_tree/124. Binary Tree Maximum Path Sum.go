package binary_tree

import (
	"leetcode/src/define/mytreenode"
	"leetcode/src/utils/mathhelper"
	"math"
)

var pathSum int

func maxPathSum(root *mytreenode.TreeNode) int {
	pathSum = math.MinInt
	maxPathSumCore(root)
	return pathSum
}

func maxPathSumCore(root *mytreenode.TreeNode) int {
	if root == nil {
		return 0
	}
	leftSum := mathhelper.Max(maxPathSumCore(root.Left), 0)
	rightSum := mathhelper.Max(maxPathSumCore(root.Right), 0)
	pathSum = mathhelper.Max(pathSum, root.Val+leftSum+rightSum)
	return mathhelper.Max(leftSum, rightSum) + root.Val
}
