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

func RemoveDuplicatesInPlace(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	set := make(map[int]bool)
	j := 0
	for i := 0; i < len(nums); i++ {
		if ok := set[nums[i]]; !ok {
			set[nums[i]] = true
			nums[j] = nums[i]
			j++
		}
	}
	return nums[:j]
}

func RemoveDuplicatesInPlaceWithOutLibraries(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	j := 1
	for i := 1; i < len(nums); i++ {
		duplicate := false
		for k := 0; k < i; k++ {
			if nums[i] == nums[k] {
				duplicate = true
				break
			}
		}
		if !duplicate {
			nums[j] = nums[i]
			j++
		}
	}
	return nums[:j]
}
