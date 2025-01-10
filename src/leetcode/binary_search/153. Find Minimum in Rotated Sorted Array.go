package binarysearch

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
// finish times: 2
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] <= nums[len(nums)-1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

func findMinV2(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] <= nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
