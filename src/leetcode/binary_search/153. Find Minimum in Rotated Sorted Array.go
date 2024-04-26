package binarysearch

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] <= nums[len(nums)-1] {
			right = mid
		} else if nums[mid] >= nums[0] {
			// when nums = [2, 1], len = 2, mid = 0, left should be left + 1
			left = left + 1
		}
	}
	return nums[left]
}
