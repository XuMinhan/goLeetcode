package main

func numTrees(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		sum := 0
		for j := 0; j <= i-1; j++ {
			sum += dp[j] * dp[i-1-j]
		}
		dp[i] = sum
	}
	return dp[n]
}
