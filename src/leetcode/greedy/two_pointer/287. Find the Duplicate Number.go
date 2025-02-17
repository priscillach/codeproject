package two_pointer

// https://leetcode.com/problems/find-the-duplicate-number/description/
func findDuplicate(nums []int) int {
	left, right := 1, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count <= mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// refer to 142. Linked List Cycle II
func findDuplicateV2(nums []int) int {
	// pre-walk once, in case escape directly from for slow != fast {}
	slow := nums[nums[0]]
	fast := nums[nums[nums[0]]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
