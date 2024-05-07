package binary_tree

import "leetcode/src/define/mytreenode"

func treeToDoublyList(root *mytreenode.TreeNode) *mytreenode.TreeNode {
	left, right := treeToDoublyListCore(root)
	left.Left = right
	right.Right = left
	return left
}

func treeToDoublyListCore(root *mytreenode.TreeNode) (*mytreenode.TreeNode, *mytreenode.TreeNode) {
	if root == nil {
		return nil, nil
	}
	left, right := root, root
	leftFront, leftTail := treeToDoublyListCore(root.Left)
	rightFront, rightTail := treeToDoublyListCore(root.Right)
	if leftFront != nil {
		left = leftFront
		leftTail.Right = root
		root.Left = leftTail
	}
	if rightFront != nil {
		right = rightTail
		rightFront.Left = root
		root.Right = rightFront
	}

	return left, right
}
