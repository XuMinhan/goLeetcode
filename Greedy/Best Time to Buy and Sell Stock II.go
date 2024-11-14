package main

func maxProfit(prices []int) int {
	l := len(prices)
	ret := 0
	for i := 1; i < l; i++ {
		if prices[i] > prices[i-1] {
			ret += prices[i] - prices[i-1]
		}
	}
	return ret
}
