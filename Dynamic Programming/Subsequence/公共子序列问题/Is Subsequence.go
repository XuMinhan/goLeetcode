package 公共子序列问题

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	if t == "" {
		return false
	}
	s1 := []byte(s)
	s2 := []byte(t)
	dp := make([][]int, len(s1))
	for i := range dp {
		dp[i] = make([]int, len(s2))
	}
	for i, n := range s1 {
		for j, m := range s2 {
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
	return dp[len(s1)-1][len(s2)-1] == len(s)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
