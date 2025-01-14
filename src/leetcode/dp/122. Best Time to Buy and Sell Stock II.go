package dp

import (
	"leetcode/src/utils/mathhelper"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/
// unlimited transaction times and only one stock hold at a time
func maxProfit(prices []int) int {
	sell, buy := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		sell = mathhelper.Max(sell, buy+prices[i])
		buy = mathhelper.Max(buy, sell-prices[i])
	}
	return sell
}

func maxProfitV2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profit := 0
	prices = append(prices, 0)
	for i := 1; i < len(prices); i++ {
		if prices[i] <= prices[i-1] {
			continue
		}
		profit += prices[i] - prices[i-1]
	}
	return profit
}
