package tool

import (
	"fmt"
	"strings"
	"testing"
)

func TestTransLCP(t *testing.T) {
	transCode := `
func deleteDuplicates(head *mylinkednode.ListNode) *mylinkednode.ListNode {
	newHead := &mylinkednode.ListNode{
		Next: head,
		Val:  math.MinInt,
	}
	cur := newHead
	for cur != nil {
		if next, isDuplicate := detectDuplicates(cur.Next); isDuplicate {
			cur.Next = next
			continue
		}
		cur = cur.Next
	}
	return newHead.Next
}

func detectDuplicates(head *mylinkednode.ListNode) (*mylinkednode.ListNode, bool) {
	if head == nil {
		return nil, false
	}
	val := head.Val
	cnt := 0
	head = head.Next
	for head != nil {
		if head.Val != val {
			break
		}
		head = head.Next
		cnt++
	}
	return head, cnt > 0
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
