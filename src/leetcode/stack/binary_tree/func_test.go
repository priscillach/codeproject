package binary_tree

import (
	"leetcode/src/define/mytreenode"
	"leetcode/src/utils/arrayhelper"
	"math"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	root := mytreenode.BuildBinaryTreeV2(arrayhelper.NumsSlice2NumsPtrSlice([]int{1, math.MinInt, 2, 3}))
	inorderTraversal(root)
}

func TestPostorderTraversal(t *testing.T) {
	root := mytreenode.BuildBinaryTreeV2(arrayhelper.NumsSlice2NumsPtrSlice([]int{1, 2, 3, 4, math.MinInt, 5, 6, math.MinInt, 7, math.MinInt, math.MinInt, 8, 9}))
	postorderTraversal(root)
}
