package greedy

func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, n+m)
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			res[i+j] = nums1[i]
			i++
		} else {
			res[i+j] = nums2[j]
			j++
		}
	}
	for i < m {
		res[i+j] = nums1[i]
		i++
	}
	for j < n {
		res[i+j] = nums2[j]
		j++
	}
	copy(nums1, res)
}
