package main

func minCostClimbingStairs(cost []int) int {
	/*
		Input: cost = [1,100,1,1,1,100,1,1,100,1]
		Output: 6
	*/
	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	if len(cost) == 1 {
		return 0
	}
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	// [10,15,20]
	//. 0  0. 10
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	dp[2] = min(cost[1], cost[0])
	for i := 3; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}
