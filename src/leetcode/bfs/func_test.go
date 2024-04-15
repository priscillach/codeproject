package bfs

import (
	"github.com/stretchr/testify/assert"
	"leetcode/src/define/mytreenode"
	"leetcode/src/utils"
	"math"
	"testing"
)

func TestWidthOfBinaryTree(t *testing.T) {
	assert.Equal(t, widthOfBinaryTree(mytreenode.BuildBinaryTreeV2(utils.NumsSlice2NumsPtrSlice([]int{1, 3, 2, 5, 3, math.MinInt, 9}))), 4)
}
