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
