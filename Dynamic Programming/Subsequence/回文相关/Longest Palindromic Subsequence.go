package main

func longestPalindromeSubseq(s string) int {
	/*
		   	Input: s = "bbbab"
		      Output: 4
		      Explanation: One possible longest palindromic subsequence is "bbbb".

			dp[i][j] = dp[i+1][j-1] + 2   n,m 相同
			dp[i][j] = max(dp[i][j-1],dp[i+1][j])   n,m 不相同
	*/
	bytes := []byte(s)
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i == j {
				dp[i][j] = 1
				continue
			}
			if j == i+1 {
				if bytes[j] == bytes[i] {
					dp[i][j] = 2
				} else {
					dp[i][j] = 1
				}
				continue
			}
			if bytes[j] == bytes[i] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][len(s)-1]

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
