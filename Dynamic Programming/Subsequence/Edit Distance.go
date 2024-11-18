package main

func minDistance(word1 string, word2 string) int {
	/*
		输入：word1 = "horse", word2 = "ros"
		输出：3
		解释：
		horse -> rorse (将 'h' 替换为 'r')
		rorse -> rose (删除 'r')
		rose -> ros (删除 'e')
	*/
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1); i++ {
		dp[i+1][0] = i + 1
	}
	for i := 0; i < len(word2); i++ {
		dp[0][i+1] = i + 1
	}
	for i, m := range word1 {
		for j, n := range word2 {
			if m == n {
				dp[i+1][j+1] = dp[i][j]
				continue
			}
			dp[i+1][j+1] = min(min(dp[i][j+1], dp[i+1][j])+1, dp[i][j]+1)
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
