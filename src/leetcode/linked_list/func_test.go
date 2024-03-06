package linked_list

import (
	"fmt"
	"leetcode/src/common/linkednode"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	list1 := &linkednode.ListNode{
		Val: 1,
		Next: &linkednode.ListNode{
			Val: 3,
			Next: &linkednode.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list2 := &linkednode.ListNode{
		Val: 1,
		Next: &linkednode.ListNode{
			Val: 2,
			Next: &linkednode.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	fmt.Println(mergeTwoLists(list1, list2))
	fmt.Println(mergeTwoListsV2(list1, list2))
}
