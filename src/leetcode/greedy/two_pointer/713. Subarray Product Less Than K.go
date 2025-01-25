package two_pointer

// https://leetcode.com/problems/subarray-product-less-than-k/description/
// finish times: 1
func numSubarrayProductLessThanK(nums []int, k int) int {
	var cnt int
	p := 0
	product := 1
	for q := 0; q < len(nums); q++ {
		product *= nums[q]
		for product >= k && p < q {
			product /= nums[p]
			p++
		}
		if product < k {
			// 以q结尾的满足product < k的subarray有q - p + 1个
			cnt += q - p + 1
		}
	}
	return cnt
}
