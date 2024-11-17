package main

func maxProfit4(k int, prices []int) int {
	dp := make([]int, 2*k)
	for i, price := range prices {
		if i == 0 {
			for j := range dp {
				if j%2 == 0 {
					dp[j] = -price
				}
			}
			continue
		}
		for j := range dp {
			if j == 0 {
				dp[j] = max(dp[j], -price)
				continue
			}
			if j%2 == 1 {
				dp[j] = max(dp[j-1]+price, dp[j])
			} else {
				dp[j] = max(dp[j-1]-price, dp[j])
			}
		}
	}
	return dp[2*k-1]
}
