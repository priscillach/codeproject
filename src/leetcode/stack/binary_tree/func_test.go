package binary_tree

import (
	"fmt"
	"leetcode/src/define/mytreenode"
	"leetcode/src/utils/arrayhelper"
	"math"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	root := mytreenode.BuildBinaryTreeV2(arrayhelper.NumsSlice2NumsPtrSlice([]int{1, math.MinInt, 2, 3}))
	fmt.Println(inorderTraversal(root))
}

func TestPostorderTraversal(t *testing.T) {
	root := mytreenode.BuildBinaryTreeV2(arrayhelper.NumsSlice2NumsPtrSlice([]int{1, 2, 3, 4, math.MinInt, 5, 6, math.MinInt, 7, math.MinInt, math.MinInt, 8, 9}))
	root = mytreenode.BuildBinaryTreeFromLeetCodeCase("[1,2,3,4,5,null,8,null,null,6,7,9]")
	fmt.Println(postorderTraversal(root))
}
