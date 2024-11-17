package main

import "math"

func maxProfit10(prices []int) int {

	ret := 0
	minNum := math.MaxInt
	for _, price := range prices {
		if price-minNum > ret {
			ret = price - minNum
		}
		if price < minNum {
			minNum = price
		}
	}
	return ret
}
func maxProfit11(prices []int) int {
	dp := make([]int, 2)
	for i, price := range prices {
		if i == 0 {
			dp[i] = -prices[i]
			continue
		}
		dp[0] = max(dp[0], -price)
		dp[1] = max(dp[1], price+dp[0])
	}
	return dp[1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
