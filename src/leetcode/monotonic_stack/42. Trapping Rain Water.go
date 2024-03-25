package monotonic_stack

import "leetcode/src/utils"

func trap(height []int) (ans int) {
	n := len(height)
	if n == 0 {
		return
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = utils.Max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = utils.Max(rightMax[i+1], height[i])
	}

	for i, h := range height {
		ans += utils.Min(leftMax[i], rightMax[i]) - h
	}
	return
}

func trapV2(height []int) (ans int) {
	var stack []int
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := utils.Min(height[left], h) - height[top]
			ans += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return
}
