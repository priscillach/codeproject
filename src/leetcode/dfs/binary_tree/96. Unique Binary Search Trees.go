package binary_tree

var memBST map[int]int

func numTrees(n int) int {
	memBST = make(map[int]int)
	return numTreesDfs(0, n-1)
}

func numTreesDfs(left, right int) int {
	if left == right {
		return 1
	}
	if v, ok := memBST[right-left+1]; ok {
		return v
	}
	sumBST := 0
	for i := left; i <= right; i++ {
		if i == left {
			leftSum := numTreesDfs(left-1, left-1)
			rightSum := numTreesDfs(left+1, right)
			sumBST += leftSum * rightSum
		} else if i == right {
			leftSum := numTreesDfs(left, right-1)
			rightSum := numTreesDfs(right+1, right+1)
			sumBST += leftSum * rightSum
		} else {
			leftSum := numTreesDfs(left, i-1)
			rightSum := numTreesDfs(i+1, right)
			sumBST += leftSum * rightSum
		}
	}
	memBST[right-left+1] = sumBST
	return sumBST
}

func numTreesV2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
