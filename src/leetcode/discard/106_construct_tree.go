package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//var idxMap map[int]int
//var postOrder, inOrder []int
//var lastPos int
//
//func buildTree(inorder []int, postorder []int) *TreeNode {
//	idxMap = make(map[int]int)
//	for idx, value := range inorder {
//		idxMap[value] = idx
//	}
//	postOrder = postorder
//	inOrder = inorder
//	lastPos = len(inorder)
//	return dfsBuildTree(0, len(inorder)-1)
//}
//
//func dfsBuildTree(left, right int) *TreeNode {
//	lastPos--
//	if left == right {
//		return &TreeNode{
//			Val: inOrder[left],
//		}
//	}
//	if left > right {
//		return nil
//	}
//	midIdx := idxMap[postOrder[lastPos]]
//	rightNode := dfsBuildTree(midIdx+1, right)
//	leftNode := dfsBuildTree(left, midIdx-1)
//	return &TreeNode{
//		Val:   inOrder[midIdx],
//		Left:  leftNode,
//		Right: rightNode,
//	}
//}

var idxMap = make(map[int]int)
var preOrder, inOrder []int
var firstPos = -1

func buildTree2(preorder []int, inorder []int) *TreeNode {
	for idx, val := range inorder {
		idxMap[val] = idx
	}
	preOrder = preorder
	inOrder = inorder
	a := dfsBuildTree2(0, len(inOrder)-1)
	fmt.Printf("%v", a)
	return a
}

func dfsBuildTree2(left, right int) *TreeNode {
	if left > right {
		return nil
	}
	firstPos++
	if left == right {
		return &TreeNode{
			Val: inOrder[left],
		}
	}
	midIdx := idxMap[preOrder[firstPos]]
	leftNode := dfsBuildTree2(left, midIdx-1)
	rightNode := dfsBuildTree2(midIdx+1, right)
	return &TreeNode{
		Val:   inOrder[midIdx],
		Left:  leftNode,
		Right: rightNode,
	}
}
