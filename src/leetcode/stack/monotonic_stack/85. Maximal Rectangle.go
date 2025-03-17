package monotonic_stack

import "leetcode/src/utils/mathhelper"

func maximalRectangle(matrix [][]byte) int {
	maxArea := 0
	m, n := len(matrix), len(matrix[0])
	height := make([]int, n)
	for i := 0; i < m; i++ {
		left, right := make([]int, n), make([]int, n)
		for j := 0; j < n; j++ {
			left[j] = -1
			right[j] = n

			if matrix[i][j] == '1' {
				height[j] += 1
			} else {
				height[j] = 0
			}
		}
		var stack []int
		for j := 0; j < n; j++ {
			for len(stack) > 0 && height[stack[len(stack)-1]] > height[j] {
				right[stack[len(stack)-1]] = j
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, j)
		}

		stack = []int{}
		for j := n - 1; j >= 0; j-- {
			for len(stack) > 0 && height[stack[len(stack)-1]] > height[j] {
				left[stack[len(stack)-1]] = j
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, j)
		}

		for j := 0; j < n; j++ {
			maxArea = mathhelper.Max(maxArea, height[j]*(right[j]-left[j]-1))
		}
	}
	return maxArea
}
