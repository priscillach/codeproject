package linked_list

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
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
	var s HandleMassUploadSalesInfoMigrationReq
	json.Unmarshal([]byte{123, 34, 108, 105, 110, 107, 34, 58, 34, 104, 116, 116, 112, 58, 47, 47, 112, 109, 115, 45, 99, 116, 108, 45, 115, 103, 45, 48, 51, 46, 117, 115, 115, 46, 115, 104, 111, 112, 101, 101, 46, 105, 111, 47, 115, 104, 111, 112, 101, 101, 95, 115, 114, 109, 95, 108, 105, 118, 101, 47, 114, 101, 116, 97, 105, 108, 47, 56, 97, 102, 51, 48, 56, 53, 102, 45, 102, 53, 101, 102, 45, 52, 56, 99, 54, 45, 97, 99, 54, 52, 45, 100, 52, 56, 49, 57, 52, 98, 48, 102, 52, 102, 98, 34, 44, 34, 116, 97, 115, 107, 95, 105, 100, 34, 58, 34, 50, 48, 50, 52, 48, 51, 49, 53, 48, 57, 52, 54, 52, 53, 95, 95, 50, 53, 95, 51, 53, 53, 34, 44, 34, 111, 112, 101, 114, 97, 116, 111, 114, 34, 58, 34, 112, 105, 112, 105, 46, 99, 104, 101, 110, 64, 115, 104, 111, 112, 101, 101, 46, 99, 111, 109, 34, 125}, &s)
	fmt.Printf("%+v", s)
}

type HandleMassUploadSalesInfoMigrationReq struct {
	// 上传链接
	Link string `json:"link"`
	// 异步任务id
	TaskID string `json:"task_id"`
	// 操作员
	Operator string `json:"operator"`
}
