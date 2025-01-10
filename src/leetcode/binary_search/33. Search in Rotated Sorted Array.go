package binarysearch

// https://leetcode.com/problems/search-in-rotated-sorted-array/description/
// finish times: 2
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left+1)>>1
		// mid := left + (right-left)>>1 both ok
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[right] && target <= nums[right] {
			left = mid + 1
		} else if nums[mid] <= nums[right] && target > nums[right] {
			right = mid - 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
