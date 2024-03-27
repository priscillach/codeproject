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
