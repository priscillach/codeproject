package mytreenode

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
// Deprecated
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

// BinaryTree2IntPointerArr tree -> arr rule: https://support.leetcode.com/hc/en-us/articles/360011883654-What-does-1-null-2-3-mean-in-binary-tree-representation
func BinaryTree2IntPointerArr(root *TreeNode) []*int {
	var queue []*TreeNode
	var res []*int
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			pop := queue[0]
			queue = queue[1:]
			if pop != nil {
				res = append(res, &pop.Val)
				queue = append(queue, pop.Left)
				queue = append(queue, pop.Right)
			} else {
				res = append(res, nil)
			}
		}
	}
	for len(res) > 0 && res[len(res)-1] == nil {
		res = res[:len(res)-1]
	}
	return res
}

// import
func BuildBinaryTreeFromLeetCodeCase(str string) *TreeNode {
	str = strings.TrimLeft(str, "[")
	str = strings.TrimRight(str, "]")
	if str == "" {
		return nil
	}
	strs := strings.Split(str, ",")
	var nums []*int
	for i := 0; i < len(strs); i++ {
		if strs[i] == LeetCodeNilTreeNode {
			nums = append(nums, nil)
		} else {
			num, err := strconv.Atoi(strs[i])
			if err != nil {
				panic(fmt.Sprintf("Split err: %v", err))
			}
			nums = append(nums, &num)
		}
	}
	return BuildBinaryTreeV2(nums)
}

// import
func BinaryTree2LeetCodeCase(root *TreeNode) string {
	nums := BinaryTree2IntPointerArr(root)
	var res []string
	for i := 0; i < len(nums); i++ {
		if nums[i] == nil {
			res = append(res, LeetCodeNilTreeNode)
		} else {
			res = append(res, strconv.Itoa(*nums[i]))
		}
	}
	return "[" + strings.Join(res, ",") + "]"
}
