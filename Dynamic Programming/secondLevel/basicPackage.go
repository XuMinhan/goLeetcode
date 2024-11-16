package main

func compute(pack, weights []int, space int) int {
	dp := make([]int, space+1)
	for index, weight := range weights {
		for i := space; i >= weight; i-- {
			dp[i] = max(dp[i-weight]+pack[index], dp[i])
		}
	}
	return dp[space]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
