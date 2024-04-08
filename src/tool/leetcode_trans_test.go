package tool

import (
	"fmt"
	"strings"
	"testing"
)

func TestTransLCP(t *testing.T) {
	transCode := `
var diameterOfBinaryTreeMax int
func diameterOfBinaryTree(root *mytreenode.TreeNode) int {
	diameterOfBinaryTreeMax = 0
	diameterOfBinaryTreeCore(root)
	return diameterOfBinaryTreeMax
}

func diameterOfBinaryTreeCore(root *mytreenode.TreeNode) int {
	if root == nil {
		return 0
	}
	depthLeft := diameterOfBinaryTree(root.Left)
	depthRight := diameterOfBinaryTree(root.Right)
	diameterOfBinaryTreeMax = utils.Max(diameterOfBinaryTreeMax, depthLeft + depthRight)
	return utils.Max(depthLeft, depthRight) + 1
}
`
	transCode = strings.ReplaceAll(transCode, "mytreenode.", "")
	transCode = strings.ReplaceAll(transCode, "mylinkednode.", "")
	if strings.Contains(transCode, "utils.Max") {
		transCode = strings.ReplaceAll(transCode, "utils.Max", "max")
		transCode += `
func max(nums... int) int {
	max := math.MinInt32
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
`
	}
	if strings.Contains(transCode, "utils.Min") {
		transCode = strings.ReplaceAll(transCode, "utils.Min", "min")
		transCode += `
func min(nums... int) int {
	min := math.MaxInt32
	for i := range nums {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
`
	}

	if strings.Contains(transCode, "utils.Abs") {
		transCode = strings.ReplaceAll(transCode, "utils.Abs", "abs")
		transCode += `
func Abs(num int) int {
	return int(math.Abs(float64(num)))
}
`
	}
	fmt.Println(transCode)
}
