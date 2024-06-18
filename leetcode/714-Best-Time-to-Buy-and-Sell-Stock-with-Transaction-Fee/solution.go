package leetcode

import "math"

func maxProfit(prices []int, fee int) int {
	notHoldingMaxProfit, holdingMaxProfit := 0, math.MinInt32

	for _, price := range prices {
		prevNotHolding := notHoldingMaxProfit
		notHoldingMaxProfit = max(notHoldingMaxProfit, holdingMaxProfit+price)
		holdingMaxProfit = max(holdingMaxProfit, prevNotHolding-price-fee)
	}
	return notHoldingMaxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
