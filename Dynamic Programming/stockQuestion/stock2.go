package main

func maxProfit20(prices []int) int {
	l := len(prices)
	ret := 0
	for i := 1; i < l; i++ {
		if prices[i] > prices[i-1] {
			ret += prices[i] - prices[i-1]
		}
	}
	return ret
}
func maxProfit21(prices []int) int {
	dp := make([]int, 2)
	for i, price := range prices {
		if i == 0 {
			dp[0] = -price
			continue
		}
		dp[0] = max(dp[0], dp[1]-price)
		dp[1] = max(dp[1], dp[0]+price)
	}
	return dp[1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
