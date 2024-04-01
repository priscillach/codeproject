package unclassified

import "math"

func firstMissingPositive(nums []int) int {
	// assume nums 1 2 .. len(nums)
	// nums append, make len(nums) can find the corresponding index
	nums = append(nums, math.MinInt)

	for idx, num := range nums {
		for idx != num {
			if num >= len(nums) || num <= 0 {
				break
			}
			tmp := nums[num]
			nums[num] = num
			idx = num
			num = tmp
		}
	}

	for i := 1; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return len(nums)
}
