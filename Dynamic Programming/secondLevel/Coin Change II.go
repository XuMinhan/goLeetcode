package main

func change(amount int, coins []int) int {
	/*
	   	Input: amount = 5, coins = [1,2,5]
	      Output: 4
	      Explanation: there are four ways to make up the amount:
	      5=5
	      5=2+2+1
	      5=2+1+1+1
	      5=1+1+1+1+1
	*/
	//0 1 2 3 4 5
	//1 0 0 0 0 0
	//1 1 1 1 1 1
	//1 1 2 2 3 3

	// dp[i][n] 代表 i号硬币加入后，组合为 n 元的组合方式
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = dp[j-coin] + dp[j]
		}
	}
	return dp[amount]
}
