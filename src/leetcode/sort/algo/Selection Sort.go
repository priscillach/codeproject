package algo

import "math"

// SelectionSort
func SelectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minNum := math.MaxInt
		var minIdx int
		for j := i; j < len(nums); j++ {
			if nums[j] < minNum {
				minNum = nums[j]
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}
