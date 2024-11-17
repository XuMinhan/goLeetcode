package main

func maxProfit30(prices []int) int {
	state1, state2, state3, state4 := 0, 0, 0, 0
	for i, price := range prices {
		if i == 0 {
			state1 = -price
			state3 = -price
			continue
		}
		state1 = max(-price, state1)
		state2 = max(price+state1, state2)
		state3 = max(state3, state2-price)
		state4 = max(state4, state3+price)
	}
	return state4

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
