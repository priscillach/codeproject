package bfs

import (
	"github.com/stretchr/testify/assert"
	"leetcode/src/define/mytreenode"
	"leetcode/src/utils/arrayhelper"
	"math"
	"testing"
)

func TestWidthOfBinaryTree(t *testing.T) {
	assert.Equal(t, widthOfBinaryTree(mytreenode.BuildBinaryTreeV2(arrayhelper.NumsSlice2NumsPtrSlice([]int{1, 3, 2, 5, 3, math.MinInt, 9}))), 4)
}

func TestSerializeBinaryTree(t *testing.T) {
	root := mytreenode.BuildBinaryTreeV2(arrayhelper.NumsSlice2NumsPtrSlice([]int{1, 2, 3, math.MinInt, math.MinInt, 4, 5}))
	ser := Constructor()
	deser := Constructor()
	data := ser.serializeOnlyForUniqueValue(root)
	ans := deser.deserializeOnlyForUniqueValue(data)
	assert.Equal(t, ans, root)

	root = mytreenode.BuildBinaryTreeV2(arrayhelper.NumsSlice2NumsPtrSlice([]int{1, 2, 3, math.MinInt, math.MinInt, 4, 5}))
	ser = Constructor()
	deser = Constructor()
	data = ser.serialize(root)
	ans = deser.deserialize(data)
	assert.Equal(t, ans, root)

	root = mytreenode.BuildBinaryTreeV2(arrayhelper.NumsSlice2NumsPtrSlice([]int{4, -7, -3, math.MinInt, math.MinInt, -9, -3, 9, -7, -4, math.MinInt, 6, math.MinInt, -6, -6, math.MinInt, math.MinInt, 0, 6, 5, math.MinInt, 9, math.MinInt, math.MinInt, -1, -4, math.MinInt, math.MinInt, math.MinInt, -2}))
	ser = Constructor()
	deser = Constructor()
	data = ser.serialize(root)
	ans = deser.deserialize(data)
	assert.Equal(t, ans, root)

	root = mytreenode.BuildBinaryTreeFromLeetCodeCase("[]")
	ser = Constructor()
	deser = Constructor()
	data = ser.serialize(root)
	ans = deser.deserialize(data)
	assert.Equal(t, ans, root)

	root = mytreenode.BuildBinaryTreeFromLeetCodeCase("[1,2,3,null,null,4,5]")
	ser = Constructor()
	deser = Constructor()
	data = ser.serialize(root)
	ans = deser.deserialize(data)
	assert.Equal(t, ans, root)

	root = mytreenode.BuildBinaryTreeFromLeetCodeCase("[4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2]")
	ser = Constructor()
	deser = Constructor()
	data = ser.serialize(root)
	ans = deser.deserialize(data)
	assert.Equal(t, ans, root)
}

func TestIsCompleteTree(t *testing.T) {
	assert.Equal(t, true, isCompleteTree(mytreenode.BuildBinaryTreeFromLeetCodeCase("[1,2,3,4,5,6]")))
}
