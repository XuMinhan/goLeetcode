package 公共子序列问题

func minDistance(word1 string, word2 string) int {
	/*
	   	Input: text1 = "abcde", text2 = "ace"
	      Output: 3
	      Explanation: The longest common subsequence is "ace" and its length is 3.
	*/
	s1 := []byte(word1)
	s2 := []byte(word2)
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
	return len(s1) + len(s2) - 2*dp[len(s1)-1][len(s2)-1]
}
