package prefix_tree

import "leetcode/src/utils/mathhelper"

func findKthNumber(n int, k int) int {
	cur := 1
	k--
	for k > 0 {
		count := countSteps(n, cur, cur+1)
		if count <= k {
			k -= count
			cur++
		} else {
			k--
			cur *= 10
		}
	}
	return cur
}

func countSteps(n int, prefix1, prefix2 int) int {
	steps := 0
	for prefix1 <= n {
		steps += mathhelper.Min(n+1, prefix2) - prefix1
		prefix1 *= 10
		prefix2 *= 10
	}
	return steps
}
