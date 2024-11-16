package main

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, num := range stones {
		sum += num
	}
	s := sum
	sum = sum / 2
	dp := make([]int, sum+1)
	for _, num := range stones {
		for i := sum; i >= num; i-- {
			dp[i] = max(dp[i], dp[i-num]+num)
		}
	}
	return s - 2*dp[sum]
}
