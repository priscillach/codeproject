package in_place_replace

func findDuplicates(nums []int) []int {
	var res []int
	nums = append([]int{0}, nums...)
	for idx, num := range nums {
		if idx != num {
			nums[idx] = -1
		}
		for idx != num {
			if num == -1 {
				break
			}
			next := nums[num]
			if next == num {
				res = append(res, num)
				break
			}
			nums[num] = num
			idx = num
			num = next
		}
	}
	return res
}

func findDuplicatesV2(nums []int) []int {
	var res []int
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if nums[num-1] < 0 {
			res = append(res, num)
		} else {
			nums[num-1] = -nums[num-1]
		}
	}
	return res
}
