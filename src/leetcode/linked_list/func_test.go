package linked_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/src/define/mylinkednode"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	list1 := &mylinkednode.ListNode{
		Val: 1,
		Next: &mylinkednode.ListNode{
			Val: 3,
			Next: &mylinkednode.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list2 := &mylinkednode.ListNode{
		Val: 1,
		Next: &mylinkednode.ListNode{
			Val: 2,
			Next: &mylinkednode.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	fmt.Println(mergeTwoLists(list1, list2))
	fmt.Println(mergeTwoListsV2(list1, list2))
}

func TestReverseBetween(t *testing.T) {
	head := &mylinkednode.ListNode{
		Val: 1,
		Next: &mylinkednode.ListNode{
			Val: 2,
			Next: &mylinkednode.ListNode{
				Val: 3,
				Next: &mylinkednode.ListNode{
					Val: 4,
					Next: &mylinkednode.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	assert.Equal(t, reverseBetween(head, 2, 4).Val, 1)
}

func TestMergeKLists(t *testing.T) {
	// [[1,4,5],[1,3,4],[2,6]]
	lists := []*mylinkednode.ListNode{
		{
			Val: 1,
			Next: &mylinkednode.ListNode{
				Val: 4,
				Next: &mylinkednode.ListNode{
					Val: 5,
				},
			},
		},
		{
			Val: 1,
			Next: &mylinkednode.ListNode{
				Val: 3,
				Next: &mylinkednode.ListNode{
					Val: 4,
				},
			},
		},
		{
			Val: 2,
			Next: &mylinkednode.ListNode{
				Val: 6,
			},
		},
	}
	mergeKLists(lists)

}
