package easy子序列问题

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums)) // dp[]
	ret := nums[0]
	for i, num := range nums {
		if i == 0 {
			dp[i] = num
		} else if dp[i-1] < 0 {
			dp[i] = num
		} else {
			dp[i] = dp[i-1] + num
		}
		ret = max(ret, dp[i])

	}
	return ret
}
