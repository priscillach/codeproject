package binary_tree

var memBST map[int]int

func numTrees(n int) int {
	memBST = make(map[int]int)
	return dfs(0, n-1)
}

func dfs(left, right int) int {
	if left == right {
		return 1
	}
	if v, ok := memBST[right-left+1]; ok {
		return v
	}
	sumBST := 0
	for i := left; i <= right; i++ {
		if i == left {
			leftSum := dfs(left-1, left-1)
			rightSum := dfs(left+1, right)
			sumBST += leftSum * rightSum
		} else if i == right {
			leftSum := dfs(left, right-1)
			rightSum := dfs(right+1, right+1)
			sumBST += leftSum * rightSum
		} else {
			leftSum := dfs(left, i-1)
			rightSum := dfs(i+1, right)
			sumBST += leftSum * rightSum
		}
	}
	memBST[right-left+1] = sumBST
	return sumBST
}
