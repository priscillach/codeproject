package divide_and_conquer

import "leetcode/src/define/mytreenode"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var preOrder889, postOrder []int
var postIdxMap map[int]int

func constructFromPrePost(preorder []int, postorder []int) *mytreenode.TreeNode {
	preOrder889 = preorder
	postOrder = postorder
	postIdxMap = make(map[int]int)
	for idx, val := range postorder {
		postIdxMap[val] = idx
	}
	return dfsBuild889(0, len(preorder)-1, 0, len(preorder)-1)
}

func dfsBuild889(preLeft, preRight, postLeft, postRight int) *mytreenode.TreeNode {
	if preLeft > preRight || postLeft > postRight {
		return nil
	}

	root := &mytreenode.TreeNode{
		Val: preOrder889[preLeft],
	}
	if preLeft == preRight {
		return root
	}
	postIdx := postIdxMap[preOrder889[preLeft+1]]
	leftNum := postIdx - postLeft
	root.Left = dfsBuild889(preLeft+1, preLeft+1+leftNum, postLeft, postIdx)
	root.Right = dfsBuild889(preLeft+2+leftNum, preRight, postIdx+1, postRight-1)
	return root
}
