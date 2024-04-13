package binarysearch

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if (mid == 0 || nums[mid] > nums[mid-1]) && (mid == len(nums)-1 || nums[mid] > nums[mid+1]) {
			return mid
		} else if (mid == 0 && nums[mid] < nums[mid+1]) || (nums[mid] > nums[mid-1] && nums[mid] < nums[mid+1]) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
