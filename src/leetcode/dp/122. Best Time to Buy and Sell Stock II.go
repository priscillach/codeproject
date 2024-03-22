package dp

import "leetcode/src/utils"

func maxProfit(prices []int) int {
	cash, hold := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		cash = utils.Max(cash, hold+prices[i])
		hold = utils.Max(hold, cash-prices[i])
	}
	return cash
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
