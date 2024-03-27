package mytreenode

import (
	"errors"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildBinaryTree [-10,9,20,null,null,15,7]
func BuildBinaryTree(nums []*int) (*TreeNode, error) {
	if len(nums) == 0 || nums[0] == nil {
		return nil, nil

	}
	var treeNodes []*TreeNode
	for idx, num := range nums {
		if num == nil {
			treeNodes = append(treeNodes, nil)
			continue
		}
		treeNode := &TreeNode{
			Val: *num,
		}
		treeNodes = append(treeNodes, treeNode)
		if idx == 0 {
			continue
		}
		parentIdx := (idx - 1) / 2
		if treeNodes[parentIdx] == nil {
			return nil, errors.New("error child node cant find parent")
		}
		if idx%2 == 1 {
			treeNodes[parentIdx].Left = treeNode
		} else {
			treeNodes[parentIdx].Right = treeNode
		}
	}
	return treeNodes[0], nil
}
