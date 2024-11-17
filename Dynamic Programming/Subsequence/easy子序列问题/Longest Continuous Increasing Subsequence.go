package easy子序列问题

func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	for i, num := range nums {
		dp[i] = 1
		if i == 0 {
			continue
		}
		if num > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
	}
	ret := 0
	for _, i := range dp {
		ret = max(ret, i)
	}
	return ret
}
