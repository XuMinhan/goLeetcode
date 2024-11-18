package main

func findLength(nums1 []int, nums2 []int) int {
	/*
	   	输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
	      输出：3
	      解释：长度最长的公共子数组是 [3,2,1] 。
	*/
	/*
		  i/j 0 1 2 3 4 5 6
			0 0 0 0 0 0 0 0
			1 0 0 0 0 1 0 0
			2
			3
			4
			5
			6
	*/
	//	dp[i][j] 是 n1[:i+1]到n2[:j+1] 的后缀一样的长度
	dp := make([][]int, len(nums1))
	for i := range dp {
		dp[i] = make([]int, len(nums2))
	}
	ret := 0
	for i, n := range nums1 {
		for j, m := range nums2 {
			if n == m {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
				ret = max(ret, dp[i][j])
			}
		}
	}
	return ret
}
