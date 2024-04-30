package binarysearch

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
