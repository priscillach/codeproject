package tool

import (
	"fmt"
	"strings"
	"testing"
)

func TestTransLCP(t *testing.T) {
	transCode := `
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
`
	transCode = strings.ReplaceAll(transCode, "mytreenode.", "")
	transCode = strings.ReplaceAll(transCode, "mytreenode.", "")
	if strings.Contains(transCode, "utils.Max") {
		transCode = strings.ReplaceAll(transCode, "utils.Max", "max")
		transCode += `
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
`
	}
	if strings.Contains(transCode, "utils.Min") {
		transCode = strings.ReplaceAll(transCode, "utils.Min", "max")
		transCode += `
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
`
	}
	fmt.Println(transCode)
}
