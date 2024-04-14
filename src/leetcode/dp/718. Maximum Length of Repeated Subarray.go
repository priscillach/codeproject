package dp

func findLength(nums1 []int, nums2 []int) int {
	maxLen := 0
	dp := make([][]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		dp[i] = make([]int, len(nums2))
	}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
				if maxLen < dp[i][j] {
					maxLen = dp[i][j]
				}
			}
		}
	}
	return maxLen
}
