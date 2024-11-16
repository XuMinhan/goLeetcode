package main

func rob1(nums []int) int {
	/*
		   	输入：[2,7,9,3,1]
		      输出：12
		      解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
		           偷窃到的最高金额 = 2 + 9 + 1 = 12 。
		[2,1,1,2]
		2 0
		1 2
		3

	*/
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = nums[0]
	dp[0][1] = 0
	for i := 1; i < len(nums); i++ {
		dp[i][0] = dp[i-1][1] + nums[i]
		dp[i][1] = max(dp[i-1][0], dp[i-1][1])
	}
	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
