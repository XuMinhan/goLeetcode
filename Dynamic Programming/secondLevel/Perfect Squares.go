package main

func numSquares(n int) int {
	o := 0
	for i := 1; i <= 101; i++ {
		if i*i > n {
			o = i
		}
	}
	ints := make([]int, o-1)
	for i := 1; i < o; i++ {
		ints[i-1] = i * i
	}
	dp := make([]int, n+1)
	for _, num := range ints {
		for i := num; i <= n; i++ {
			if i == num || dp[i-num] != 0 {
				if dp[i] == 0 {
					dp[i] = dp[i-num] + 1
					continue
				}
				dp[i] = min(dp[i], dp[i-num]+1)
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
