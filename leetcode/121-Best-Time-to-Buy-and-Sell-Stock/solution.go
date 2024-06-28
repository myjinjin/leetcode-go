package leetcode

import "math"

func maxProfit(prices []int) int {
	profit := 0
	minPrice := math.MaxInt32
	for _, price := range prices {
		if price-minPrice > profit {
			profit = price - minPrice
		}
		minPrice = min(minPrice, price)
	}
	return profit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
