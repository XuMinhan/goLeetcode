package main

func maxProfit(prices []int, fee int) int {
	notHave := -prices[0]
	have := 0
	for i, price := range prices {
		if i == 0 {
			continue
		}
		notHave = max(notHave, -price+have)
		have = max(have, notHave+price-fee)
	}
	return have
}
