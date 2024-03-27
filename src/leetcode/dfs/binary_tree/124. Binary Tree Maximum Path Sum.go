package binary_tree

import (
	"leetcode/src/define/mytreenode"
	"leetcode/src/utils"
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
	leftSum := utils.Max(maxPathSumCore(root.Left), 0)
	rightSum := utils.Max(maxPathSumCore(root.Right), 0)
	pathSum = utils.Max(pathSum, root.Val+leftSum+rightSum)
	return utils.Max(leftSum, rightSum) + root.Val
}
