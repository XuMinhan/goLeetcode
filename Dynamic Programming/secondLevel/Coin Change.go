package main

// 0 1 2 3 4 5 6 7 8 9 10 11
// 0
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if dp[i-coin] != 0 || i == coin {
				if dp[i] == 0 {
					dp[i] = dp[i-coin] + 1
					continue
				}
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}
