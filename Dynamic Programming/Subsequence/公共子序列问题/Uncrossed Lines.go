package 公共子序列问题

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := range dp {
		dp[i] = make([]int, len(nums2))
	}
	for i, n := range nums1 {
		for j, m := range nums2 {
			if n != m {
				if i == 0 && j == 0 {
					dp[i][j] = 0
				} else if j == 0 {
					dp[i][j] = dp[i-1][j]
				} else if i == 0 {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				}
			} else {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			}
		}
	}
	return dp[len(nums1)-1][len(nums2)-1]
}
