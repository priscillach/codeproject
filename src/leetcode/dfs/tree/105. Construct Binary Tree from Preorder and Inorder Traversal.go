package tree

import "leetcode/src/define/treenode"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var idxMap105 map[int]int
var preOrder105, inOrder []int
var firstPos int

func buildTree105(preorder []int, inorder []int) *treenode.TreeNode {
	idxMap105 = make(map[int]int)
	for idx, val := range inorder {
		idxMap105[val] = idx
	}
	preOrder105 = preorder
	inOrder = inorder
	firstPos = -1
	a := dfsBuildTree105(0, len(inOrder)-1)

	return a
}

func dfsBuildTree105(left, right int) *treenode.TreeNode {
	if left > right {
		return nil
	}
	firstPos++
	if left == right {
		return &treenode.TreeNode{
			Val: inOrder[left],
		}
	}
	midIdx := idxMap105[preOrder105[firstPos]]
	leftNode := dfsBuildTree105(left, midIdx-1)
	rightNode := dfsBuildTree105(midIdx+1, right)
	return &treenode.TreeNode{
		Val:   inOrder[midIdx],
		Left:  leftNode,
		Right: rightNode,
	}
}
