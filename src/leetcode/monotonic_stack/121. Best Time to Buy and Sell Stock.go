package monotonic_stack

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var stack []int
	profit := 0
	for i := 0; i < len(prices); i++ {
		for len(stack) > 0 && prices[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, prices[i])
		if len(stack) > 1 {
			if stack[len(stack)-1]-stack[0] > profit {
				profit = stack[len(stack)-1] - stack[0]
			}
		}
	}

	return profit
}
