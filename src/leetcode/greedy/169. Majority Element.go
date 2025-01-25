package greedy

// https://leetcode.com/problems/majority-element/description/
func majorityElement(nums []int) int {
	var cnt, res = 1, nums[0]
	for _, num := range nums[1:] {
		if num == res {
			cnt++
			continue
		} else if cnt == 0 {
			res = num
			cnt++
		} else {
			cnt--
		}
	}
	return res
}

func majorityElementV2(nums []int) int {
	var cnt int
	var cur int
	for _, num := range nums {
		if cnt == 0 {
			cur = num
			cnt++
		} else if num != cur {
			cnt--
		} else {
			cnt++
		}

	}
	return cur
}
