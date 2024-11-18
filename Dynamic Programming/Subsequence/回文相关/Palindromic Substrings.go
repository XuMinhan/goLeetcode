package main

func countSubstrings(s string) int {
	bytes := []byte(s)
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	/*
		dp[i][j] =
	*/
	for i := len(s) - 1; i >= 0; i-- {
		for j := len(s) - 1; j >= i; j-- {
			if i == j || j == i+1 {
				dp[i][j] = bytes[i] == bytes[j]
				continue
			}
			dp[i][j] = dp[i+1][j-1] && bytes[i] == bytes[j]
		}
	}
	ret := 0
	for _, bools := range dp {
		for _, b := range bools {
			if b {
				ret++
			}
		}
	}
	return ret
}
