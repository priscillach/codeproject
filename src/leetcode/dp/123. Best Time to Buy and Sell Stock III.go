package dp

import (
	"leetcode/src/utils/mathhelper"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/description/
// only two transactions and only one stock hold at a time
func maxProfit123(prices []int) int {
	buy1, buy2, sell1, sell2 := -prices[0], -prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		buy1 = mathhelper.Max(buy1, -prices[i])
		sell1 = mathhelper.Max(sell1, buy1+prices[i])
		buy2 = mathhelper.Max(buy2, sell1-prices[i])
		sell2 = mathhelper.Max(sell2, buy2+prices[i])
	}
	return sell2
}
