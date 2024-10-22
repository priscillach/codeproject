package binary_tree

import "leetcode/src/define/mytreenode"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var idxMap105 map[int]int
var preOrder105, inOrder105 []int
var firstPos int

func BuildTree105(preorder []int, inorder []int) *mytreenode.TreeNode {
	idxMap105 = make(map[int]int)
	for idx, val := range inorder {
		idxMap105[val] = idx
	}
	preOrder105 = preorder
	inOrder105 = inorder
	firstPos = -1
	a := dfsBuildTree105(0, len(inOrder105)-1)

	return a
}

func dfsBuildTree105(left, right int) *mytreenode.TreeNode {
	if left > right {
		return nil
	}
	firstPos++
	if left == right {
		return &mytreenode.TreeNode{
			Val: inOrder105[left],
		}
	}
	midIdx := idxMap105[preOrder105[firstPos]]
	leftNode := dfsBuildTree105(left, midIdx-1)
	rightNode := dfsBuildTree105(midIdx+1, right)
	return &mytreenode.TreeNode{
		Val:   inOrder105[midIdx],
		Left:  leftNode,
		Right: rightNode,
	}
}
