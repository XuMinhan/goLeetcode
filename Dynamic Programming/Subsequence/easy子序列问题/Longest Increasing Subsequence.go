package easy子序列问题

func lengthOfLIS(nums []int) int {
	/*
		dp[i] 带有i的最大子序列长度
	*/
	dp := make([]int, len(nums))
	// 10,9,2,5,3,7,101,18
	//  1 1 1 2 1 2
	for i, num := range nums {
		if i == 0 {
			dp[i] = 1
			continue
		}
		for j := 0; j < i; j++ {
			if num > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		if dp[i] == 0 {
			dp[i] = 1
		}
	}
	ret := 0
	for _, num := range dp {
		println(num)
		if num > ret {
			ret = num
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
