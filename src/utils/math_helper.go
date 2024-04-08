package utils

import "math"

func Max(nums ...int) int {
	max := math.MinInt32
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func Min(nums ...int) int {
	min := math.MaxInt32
	for i := range nums {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

func NumsSlice2NumsPtrSlice(nums []int) []*int {
	var res []*int
	for _, num := range nums {
		if num == math.MinInt {
			res = append(res, nil)
		} else {
			numC := num
			res = append(res, &numC)
		}
	}
	return res
}

func Abs(num int) int {
	return int(math.Abs(float64(num)))
}
