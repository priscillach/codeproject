package binary_tree

import "leetcode/src/define/mytreenode"

// https://leetcode.com/problems/path-sum-ii/description/
// finish times: 2
var pathSum2Res [][]int

func pathSum2(root *mytreenode.TreeNode, targetSum int) [][]int {
	pathSum2Res = make([][]int, 0)
	pathSum2Core(root, targetSum, 0, []int{})
	return pathSum2Res
}

func pathSum2Core(root *mytreenode.TreeNode, targetSum, curSum int, path []int) {
	if root == nil {
		return
	}
	curSum += root.Val
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && curSum == targetSum {
		cpPath := make([]int, len(path))
		copy(cpPath, path)
		pathSum2Res = append(pathSum2Res, cpPath)
		return
	}
	pathSum2Core(root.Left, targetSum, curSum, path)
	pathSum2Core(root.Right, targetSum, curSum, path)
}
