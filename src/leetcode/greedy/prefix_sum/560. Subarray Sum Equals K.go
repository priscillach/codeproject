package prefix_sum

func subarraySum(nums []int, k int) int {
	preSum := make([]int, len(nums))
	preSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}
	var cnt int
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= i; j++ {
			right := preSum[i]
			var left int
			if j == 0 {
				left = 0
			} else {
				left = preSum[j-1]
			}
			if right-left == k {
				cnt++
			}
		}
	}
	return cnt
}

func subarraySumV2(nums []int, k int) int {

	m := make(map[int]int)
	// pre-store the prefix sum of empty nums
	m[0] = 1
	out, sum := 0, 0

	for _, v := range nums {
		sum += v
		// find sum1 meet sum2 - sum1 = k -> sum1 = sum2 - k
		prefix := sum - k

		if v, ok := m[prefix]; ok {
			out += v
		}
		// record sum1
		m[sum]++
	}

	return out
}
