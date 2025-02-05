package binary_tree

import "leetcode/src/define/mytreenode"

func isSubStructure(root *mytreenode.TreeNode, subRoot *mytreenode.TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}
	if subStructureEqual(root, subRoot) {
		return true
	}
	return isSubStructure(root.Left, subRoot) || isSubStructure(root.Right, subRoot)
}

func subStructureEqual(root1 *mytreenode.TreeNode, root2 *mytreenode.TreeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 != nil && root2 != nil {
		return subStructureEqual(root1.Left, root2.Left) && subStructureEqual(root1.Right, root2.Right) && root1.Val == root2.Val
	}

	return false
}
