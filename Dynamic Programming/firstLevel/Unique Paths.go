package main

func main() {
	println(uniquePaths(2, 2))
}
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for j := range dp {
		if j == 0 {
			for i := 0; i < n; i++ {
				dp[j] = append(dp[j], 1)
			}
		} else {
			for i := 0; i < n; i++ {
				if i == 0 {
					dp[j] = append(dp[j], 1)
				} else {
					dp[j] = append(dp[j], dp[j][i-1]+dp[j-1][i])
				}
			}
		}
	}
	return dp[m-1][n-1]

}
