package two_pointer

import "sort"

func triangleNumber(nums []int) int {
	sum := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == 0 {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[j] == 0 {
				continue
			}
			// 把j的坐标带上去，至少k = j时，一定nums[i] + nums[j] > nums[k]
			// 因为需要求最大的满足nums[i] + nums[j] > nums[k]的k，所以binary search find the last k
			// nums[k] < nums[i] + nums[j]，即nums[k] <= nums[i] + nums[j] - 1
			k := binarySearch(nums, j, len(nums)-1, nums[i]+nums[j]-1)
			sum += k - j
		}
	}
	return sum
}

func binarySearch(nums []int, left, right int, target int) int {
	for left < right {
		mid := left + (right-left+1)>>1
		if nums[mid] > target {
			right = mid - 1
		} else { // nums[mid] <= target
			left = mid
		}
	}
	return left
}
