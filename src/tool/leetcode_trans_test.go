package tool

import (
	"fmt"
	"strings"
	"testing"
)

func TestTransLCP(t *testing.T) {
	transCode := `
func sortList(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tmp := slow.Next
	slow.Next = nil
	head1 := sortList(head)
	head2 := sortList(tmp)
	return mergeTwoListsV2(head1, head2)
}

func mergeTwoListsV2(list1 *mylinkednode.ListNode, list2 *mylinkednode.ListNode) *mylinkednode.ListNode {
	cur1 := list1
	cur2 := list2
	head := &mylinkednode.ListNode{}
	cur3 := head
	for cur1 != nil && cur2 != nil {
		if cur1.Val > cur2.Val {
			cur3.Next = cur2
			cur2 = cur2.Next
		} else {
			cur3.Next = cur1
			cur1 = cur1.Next
		}
		cur3 = cur3.Next
	}
	if cur1 != nil {
		cur3.Next = cur1
	}
	if cur2 != nil {
		cur3.Next = cur2
	}
	return head.Next
}

`
	transCode = strings.ReplaceAll(transCode, "mytreenode.", "")
	transCode = strings.ReplaceAll(transCode, "mylinkednode.", "")
	if strings.Contains(transCode, "utils.Max") {
		transCode = strings.ReplaceAll(transCode, "utils.Max", "max")
		transCode += `
func max(nums... int) int {
	max := math.MinInt32
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
`
	}
	if strings.Contains(transCode, "utils.Min") {
		transCode = strings.ReplaceAll(transCode, "utils.Min", "min")
		transCode += `
func min(nums... int) int {
	min := math.MaxInt32
	for i := range nums {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
`
	}
	fmt.Println(transCode)
}
