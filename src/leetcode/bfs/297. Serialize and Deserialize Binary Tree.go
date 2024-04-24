package bfs

import (
	"leetcode/src/define/mytreenode"
	"leetcode/src/leetcode/dfs/binary_tree"
	"leetcode/src/leetcode/string_resolve"
	"leetcode/src/utils/stringhelper"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (my *Codec) serializeOnlyForUniqueValue(root *mytreenode.TreeNode) string {
	cur := root
	var preOrder297 []int
	var inOrder297 []int
	var stack []*mytreenode.TreeNode
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			preOrder297 = append(preOrder297, cur.Val)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		inOrder297 = append(inOrder297, cur.Val)
		cur = cur.Right
	}
	return stringhelper.IntArr2String(preOrder297) + "|" + stringhelper.IntArr2String(inOrder297)
}

// Deserializes your encoded data to tree.
func (my *Codec) deserializeOnlyForUniqueValue(data string) *mytreenode.TreeNode {
	arrs := strings.Split(data, "|")
	if len(arrs) != 2 {
		panic("err data to be deserialized")
	}
	return binary_tree.BuildTree105(stringhelper.String2IntArr(arrs[0]), stringhelper.String2IntArr(arrs[1]))
}

func (my *Codec) serialize(root *mytreenode.TreeNode) string {
	var queue []*mytreenode.TreeNode
	var res []string
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			pop := queue[0]
			queue = queue[1:]
			if pop != nil {
				res = append(res, strconv.Itoa(pop.Val))
				queue = append(queue, pop.Left)
				queue = append(queue, pop.Right)
			} else {
				res = append(res, "-")
			}
		}
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (my *Codec) deserialize(data string) *mytreenode.TreeNode {
	if data == "" || data == "-" {
		return nil
	}
	strs := strings.Split(data, ",")
	root := &mytreenode.TreeNode{
		Val: string_resolve.MyAtoi(strs[0]),
	}
	index := 1
	queue := []*mytreenode.TreeNode{root}
	for len(queue) > 0 && index < len(strs) {
		pop := queue[0]
		queue = queue[1:]

		if index < len(strs) && strs[index] != "-" {
			pop.Left = &mytreenode.TreeNode{
				Val: string_resolve.MyAtoi(strs[index]),
			}
			queue = append(queue, pop.Left)
		}
		index++
		if index < len(strs) && strs[index] != "-" {
			pop.Right = &mytreenode.TreeNode{
				Val: string_resolve.MyAtoi(strs[index]),
			}
			queue = append(queue, pop.Right)
		}
		index++
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
