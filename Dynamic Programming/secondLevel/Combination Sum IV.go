package main

/*
		dp[i][n] 是 i个数加入后，总和为n的可能数量
	   	Input: nums = [1,2,3], target = 4
		0 1 2 3 4
		1 1 1 1 1
		1 1 2


	      Output: 7
	      Explanation:
	      The possible combination ways are:
	      (1, 1, 1, 1)
	      (1, 1, 2)
	      (1, 2, 1)
	      (1, 3)
	      (2, 1, 1)
	      (2, 2)
	      (3, 1)
*/

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	// i 级楼梯
	dp[0] = 1
	// 1 2 3

	// 0 1 2 3 4
	// 1 1 1
	// 1 1 2
	// 1 1 2
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i < num {
				continue
			}
			dp[i] = dp[i-num] + dp[i]
		}
	}
	return dp[target]
}
