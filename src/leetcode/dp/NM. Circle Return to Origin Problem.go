package dp

// https://mp.weixin.qq.com/s/NZPaFsFrTybO3K3s7p7EVg
func numCycleReturnOrigin(n, k int) int {
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n)
	}

	// dp[i][j] represents after i steps to j index
	dp[0][0] = 1
	for i := 1; i <= k; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = dp[i-1][(j-1+n)%n] + dp[i-1][(j+1)%n]
		}
	}
	return dp[k][0]
}
