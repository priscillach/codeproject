package recursive

var climbStairsNums = make([]int, 46)

func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if climbStairsNums[n] != 0 {
		return climbStairsNums[n]
	}
	climbStairsNums[n] = climbStairs(n-1) + climbStairs(n-2)
	return climbStairsNums[n]
}
