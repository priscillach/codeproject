package linked_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/src/define/mylinkednode"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	list1 := mylinkednode.BuildLinkedList([]int{1, 3, 4})
	list2 := mylinkednode.BuildLinkedList([]int{1, 2, 4})
	fmt.Println(mergeTwoLists(list1, list2))
	fmt.Println(mergeTwoListsV2(list1, list2))
}

func TestReverseBetween(t *testing.T) {
	head := mylinkednode.BuildLinkedList([]int{1, 2, 3, 4, 5})
	assert.Equal(t, ReverseBetween(head, 2, 4).Val, 1)
}

func TestMergeKLists(t *testing.T) {
	// [[1,4,5],[1,3,4],[2,6]]
	lists := []*mylinkednode.ListNode{
		mylinkednode.BuildLinkedList([]int{1, 4, 5}),
		mylinkednode.BuildLinkedList([]int{1, 3, 4}),
		mylinkednode.BuildLinkedList([]int{2, 6}),
	}
	mergeKLists(lists)

}

func TestReorderList(t *testing.T) {
	head := mylinkednode.BuildLinkedList([]int{1, 2, 3, 4, 5})
	reorderList(head)
}

func TestDeleteDuplicates(t *testing.T) {
	deleteDuplicates(mylinkednode.BuildLinkedList([]int{1, 1}))
}
