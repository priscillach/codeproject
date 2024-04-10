package tool

import (
	"fmt"
	"strings"
	"testing"
)

func TestTransLCP(t *testing.T) {
	transCode := `
func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j - 1]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i - 1][j]
			} else {
				dp[i][j] = utils.Min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
			}
		}
	}
	return dp[len(grid) - 1][len(grid[0]) - 1]
}

`
	transCode = strings.ReplaceAll(transCode, "mytreenode.", "")
	transCode = strings.ReplaceAll(transCode, "mylinkednode.", "")
	if strings.Contains(transCode, "utils.Max") {
		transCode = strings.ReplaceAll(transCode, "utils.Max", "max")
		transCode += `
func max(nums... int) int {
	max := math.MinInt32
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
`
	}
	if strings.Contains(transCode, "utils.Min") {
		transCode = strings.ReplaceAll(transCode, "utils.Min", "min")
		transCode += `
func min(nums... int) int {
	min := math.MaxInt32
	for i := range nums {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
`
	}

	if strings.Contains(transCode, "utils.Abs") {
		transCode = strings.ReplaceAll(transCode, "utils.Abs", "abs")
		transCode += `
func Abs(num int) int {
	return int(math.Abs(float64(num)))
}
`
	}
	fmt.Println(transCode)
}
