package main

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < target && target > 0 {
		return 0
	}
	if target < 0 && -sum > target {
		return 0
	}
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2*sum+1)
	}
	for i, num := range nums {
		if i == 0 {
			if num == 0 {
				dp[i][sum] = 2
				continue
			}
			dp[i][num+sum] = 1
			dp[i][sum-num] = 1
			continue
		}
		for pos := 2 * sum; pos >= 0; pos-- {
			if pos-num >= 0 && pos+num <= 2*sum {
				dp[i][pos] = dp[i-1][pos-num] + dp[i-1][pos+num]
			} else if pos-num < 0 {
				dp[i][pos] = dp[i-1][pos+num]
			} else {
				dp[i][pos] = dp[i-1][pos-num]
			}
		}
	}
	return dp[len(nums)-1][target+sum]
}
