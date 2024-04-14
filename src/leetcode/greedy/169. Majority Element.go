package greedy

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
