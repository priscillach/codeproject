package arrayhelper

import "math"

func FillSlice(nums []int, num int) {
	for i := 0; i < len(nums); i++ {
		nums[i] = num
	}
}

func NumsSlice2NumsPtrSlice(nums []int) []*int {
	var res []*int
	for _, num := range nums {
		if num == math.MinInt || num == math.MinInt32 || num == math.MinInt64 {
			res = append(res, nil)
		} else {
			numC := num
			res = append(res, &numC)
		}
	}
	return res
}
