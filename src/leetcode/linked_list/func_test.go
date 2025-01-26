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
	deleteDuplicates82(mylinkednode.BuildLinkedList([]int{1, 1}))
	deleteDuplicates83(mylinkednode.BuildLinkedList([]int{1, 1}))
	deleteDuplicates82V2(mylinkednode.BuildLinkedList([]int{1, 1}))
}

func TestSortList(t *testing.T) {
	sortList(mylinkednode.BuildLinkedList([]int{4, 2, 1, 3}))
}

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, isPalindrome(mylinkednode.BuildLinkedList([]int{1, 0, 1})), true)
	isPalindromeV2(mylinkednode.BuildLinkedList([]int{1, 2, 2, 1}))
}

func TestSortOddAscEvenDescList(t *testing.T) {
	assert.Equal(t, mylinkednode.BuildLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8}), sortOddAscEvenDescList(mylinkednode.BuildLinkedList([]int{1, 8, 3, 6, 5, 4, 7, 2})))
	assert.Equal(t, mylinkednode.BuildLinkedList([]int{1, 2}), sortOddAscEvenDescList(mylinkednode.BuildLinkedList([]int{1, 2})))
	assert.Equal(t, mylinkednode.BuildLinkedList([]int{1}), sortOddAscEvenDescList(mylinkednode.BuildLinkedList([]int{1})))
	assert.Equal(t, mylinkednode.BuildLinkedList([]int{1, 2, 3, 4}), sortOddAscEvenDescList(mylinkednode.BuildLinkedList([]int{1, 4, 3, 2})))
	assert.Equal(t, mylinkednode.BuildLinkedList([]int{1, 2, 3}), sortOddAscEvenDescList(mylinkednode.BuildLinkedList([]int{1, 2, 3})))
}

func TestTrainingPlan(t *testing.T) {
	node := trainingPlan(mylinkednode.BuildLinkedList([]int{1, 2, 3, 4, 5}), 2)
	fmt.Println(node.Val)
}

func TestFindMiddleElement(t *testing.T) {
	fmt.Println(findLeftMiddleElement(mylinkednode.BuildLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8})).Val)
	fmt.Println(findLeftMiddleElement(mylinkednode.BuildLinkedList([]int{})))
	fmt.Println(findLeftMiddleElement(mylinkednode.BuildLinkedList([]int{1})).Val)
	fmt.Println(findLeftMiddleElement(mylinkednode.BuildLinkedList([]int{1, 2})).Val)
	fmt.Println(findLeftMiddleElement(mylinkednode.BuildLinkedList([]int{1, 2, 3})).Val)

	fmt.Println(findRightMiddleElement(mylinkednode.BuildLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8})).Val)
	fmt.Println(findRightMiddleElement(mylinkednode.BuildLinkedList([]int{})))
	fmt.Println(findRightMiddleElement(mylinkednode.BuildLinkedList([]int{1})).Val)
	fmt.Println(findRightMiddleElement(mylinkednode.BuildLinkedList([]int{1, 2})).Val)
	fmt.Println(findRightMiddleElement(mylinkednode.BuildLinkedList([]int{1, 2, 3})).Val)
}

func TestAddTwoNumbers(t *testing.T) {
	addTwoNumbers(mylinkednode.BuildLinkedList([]int{9, 9, 9, 9, 9, 9, 9}), mylinkednode.BuildLinkedList([]int{9, 9, 9, 9}))
}

func TestReverseKGroup(t *testing.T) {
	reverseKGroup(mylinkednode.BuildLinkedList([]int{1, 2, 3, 4, 5}), 2)
}

func TestGetIntersectionNode(t *testing.T) {
	getIntersectionNode(mylinkednode.BuildLinkedList([]int{1, 5}), mylinkednode.BuildLinkedList([]int{2, 6, 4}))
}
