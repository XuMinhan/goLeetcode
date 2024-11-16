package main

func findMaxForm(strs []string, m int, n int) int {
	source := make([][2]int, len(strs))
	for i, str := range strs {
		for _, char := range str {
			if char == '1' {
				source[i][1]++
			} else {
				source[i][0]++
			}
		}
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, ints := range source {
		for i := m; i >= ints[0]; i-- {
			for j := n; j >= ints[1]; j-- {
				dp[i][j] = max(dp[i][j], dp[i-ints[0]][j-ints[1]]+1)
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
