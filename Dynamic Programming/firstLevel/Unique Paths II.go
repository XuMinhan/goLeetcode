package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for j := range dp {
		if j == 0 {
			for i := 0; i < n; i++ {
				if obstacleGrid[j][i] == 1 {
					dp[j] = append(dp[j], 0)
					continue
				}
				if i >= 1 {
					dp[j] = append(dp[j], dp[j][i-1])
				} else {
					dp[j] = append(dp[j], 1)
				}
			}
		} else {
			for i := 0; i < n; i++ {
				if i == 0 {
					if obstacleGrid[j][i] == 1 {
						dp[j] = append(dp[j], 0)
						continue
					}
					if j >= 1 {
						dp[j] = append(dp[j], dp[j-1][i])
						continue
					}
					dp[j] = append(dp[j], 1)
				} else {
					if obstacleGrid[j][i] == 1 {
						dp[j] = append(dp[j], 0)
						continue
					}
					dp[j] = append(dp[j], dp[j][i-1]+dp[j-1][i])
				}
			}
		}
	}
	return dp[m-1][n-1]
}
