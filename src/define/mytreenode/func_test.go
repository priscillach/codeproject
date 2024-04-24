package mytreenode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/src/utils/arrayhelper"
	"math"
	"testing"
)

func TestBuildBinaryTree(t *testing.T) {
	root, err := BuildBinaryTree(arrayhelper.NumsSlice2NumsPtrSlice([]int{-10, 9, 20, math.MinInt, math.MinInt, 15, 7}))

	assert.NoError(t, err)
	assert.Equal(t, root.Right.Right.Val, 7)

	root = BuildBinaryTreeV2(arrayhelper.NumsSlice2NumsPtrSlice([]int{-10, 9, 20, math.MinInt, math.MinInt, 15, 7}))
	assert.Equal(t, root.Right.Right.Val, 7)

	_, err = BuildBinaryTree(arrayhelper.NumsSlice2NumsPtrSlice([]int{-10, 9, math.MinInt, math.MinInt, math.MinInt, 15, 7}))
	assert.Error(t, err)

	root, err = BuildBinaryTree(arrayhelper.NumsSlice2NumsPtrSlice([]int{math.MinInt}))
	assert.Nil(t, root)

	root, err = BuildBinaryTree(arrayhelper.NumsSlice2NumsPtrSlice([]int{1, math.MinInt, 2, 3}))
	assert.Error(t, err)
	assert.Nil(t, root)
	root = BuildBinaryTreeV2(arrayhelper.NumsSlice2NumsPtrSlice([]int{1, math.MinInt, 2, 3}))
	assert.Equal(t, root.Right.Left.Val, 3)

}

func TestBinaryTree2Nums(t *testing.T) {
	originalArr := arrayhelper.NumsSlice2NumsPtrSlice([]int{4, -7, -3, math.MinInt, math.MinInt, -9, -3, 9, -7, -4, math.MinInt, 6, math.MinInt, -6, -6, math.MinInt, math.MinInt, 0, 6, 5, math.MinInt, 9, math.MinInt, math.MinInt, -1, -4, math.MinInt, math.MinInt, math.MinInt, -2})
	root := BuildBinaryTreeV2(originalArr)
	decodeArr := BinaryTree2IntPointerArr(root)
	assert.Equal(t, originalArr, decodeArr)
}

func TestLeetCodeBinaryTree(t *testing.T) {
	originalStr := "[]"
	root := BuildBinaryTreeFromLeetCodeCase(originalStr)
	decodeStr := BinaryTree2LeetCodeCase(root)
	fmt.Println(originalStr)
	fmt.Println(decodeStr)
	assert.Equal(t, decodeStr, originalStr)

	originalStr = "[4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2]"
	root = BuildBinaryTreeFromLeetCodeCase(originalStr)
	decodeStr = BinaryTree2LeetCodeCase(root)
	fmt.Println(originalStr)
	fmt.Println(decodeStr)
	assert.Equal(t, decodeStr, originalStr)

	originalArr := arrayhelper.NumsSlice2NumsPtrSlice([]int{4, -7, -3, math.MinInt, math.MinInt, -9, -3, 9, -7, -4, math.MinInt, 6, math.MinInt, -6, -6, math.MinInt, math.MinInt, 0, 6, 5, math.MinInt, 9, math.MinInt, math.MinInt, -1, -4, math.MinInt, math.MinInt, math.MinInt, -2})
	root2 := BuildBinaryTreeV2(originalArr)
	assert.Equal(t, root, root2)
}
