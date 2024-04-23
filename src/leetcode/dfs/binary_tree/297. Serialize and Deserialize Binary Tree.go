package binary_tree

import "leetcode/src/define/mytreenode"

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *mytreenode.TreeNode) string {
	// todo
	return ""
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *mytreenode.TreeNode {
	// todo
	return &mytreenode.TreeNode{}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
