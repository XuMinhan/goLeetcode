package main

func numDistinct(s string, t string) int {
	/*	{}a
		{}b
		Input: s = "fsxzrabbbit", t = "rabbit"
		Output: 3
		Explanation:
		As shown below, there are 3 ways you can generate "rabbit" from s.
		dp[i][j] = dp[i-1][j-1] + dp[i-1][j] 相等		s[:i+1] 对于 t[:j+1] 最长的子序列
				 = dp[i-1][j] 不等
	*/
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(t))
	}
	for i, m := range s {
		for j, n := range t {
			if i == 0 {
				if j == 0 && m == n {
					dp[i][j] = 1
				}
				continue
			}
			if j == 0 {
				if m == n {
					dp[i][j] = dp[i-1][j] + 1
				} else {
					dp[i][j] = dp[i-1][j]
				}
				continue
			}
			if m != n {
				dp[i][j] = dp[i-1][j]
				continue
			}
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
	return dp[len(s)-1][len(t)-1]

}
