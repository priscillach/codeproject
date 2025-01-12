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

var idxMap106 map[int]int
var postOrder106, inOrder106 []int

func buildTree106(inorder []int, postorder []int) *mytreenode.TreeNode {
	idxMap106 = make(map[int]int)
	for idx, value := range inorder {
		idxMap106[value] = idx
	}
	postOrder106 = postorder
	inOrder106 = inorder
	return dfsBuildTree106(0, len(inorder)-1, 0, len(inorder)-1)
}

func dfsBuildTree106(inLeft, inRight, postLeft, postRight int) *mytreenode.TreeNode {
	if inLeft > inRight || postLeft > postRight {
		return nil
	}
	root := &mytreenode.TreeNode{
		Val: postOrder106[postRight],
	}
	if inLeft == inRight {
		return root
	}

	midIdx := idxMap106[postOrder106[postRight]]

	leftNum := midIdx - inLeft
	root.Left = dfsBuildTree106(inLeft, midIdx-1, postLeft, postLeft+leftNum-1)
	root.Right = dfsBuildTree106(midIdx+1, inRight, postLeft+leftNum, postRight-1)
	return root
}

var idxV2 int
var inorderMap map[int]int

func buildTreeV2(inorder []int, postorder []int) *mytreenode.TreeNode {
	inorderMap = make(map[int]int)
	for pos, node := range inorder {
		inorderMap[node] = pos
	}
	idxV2 = len(inorder)
	return buildTreeCore(inorder, postorder, 0, len(postorder)-1)
}

func buildTreeCore(inorder []int, postorder []int, left, right int) *mytreenode.TreeNode {
	if left > right {
		return nil
	}
	idxV2--
	if left == right {
		return &mytreenode.TreeNode{Val: inorder[left]}
	}
	val := postorder[idxV2]
	node := &mytreenode.TreeNode{Val: val}
	node.Right = buildTreeCore(inorder, postorder, inorderMap[val]+1, right)
	node.Left = buildTreeCore(inorder, postorder, left, inorderMap[val]-1)
	return node
}
