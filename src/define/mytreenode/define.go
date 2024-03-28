package mytreenode

import (
	"errors"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildBinaryTreeV2 [1, nil, 2, 3]
// BFS for every layer the nil leaf is ignored
func BuildBinaryTreeV2(nums []*int) *TreeNode {
	if len(nums) == 0 || nums[0] == nil {
		return nil
	}

	// Create the root node
	root := &TreeNode{Val: *nums[0]}
	queue := []*TreeNode{root}
	index := 1

	for len(queue) > 0 && index < len(nums) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if index < len(nums) && nums[index] != nil {
			node.Left = &TreeNode{Val: *nums[index]}
			queue = append(queue, node.Left)
		}
		index++

		// Right child
		if index < len(nums) && nums[index] != nil {
			node.Right = &TreeNode{Val: *nums[index]}
			queue = append(queue, node.Right)
		}
		index++
	}

	return root
}

// BuildBinaryTree [-10,9,20,null,null,15,7]
// for every layer the nil leaf should be filled
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
