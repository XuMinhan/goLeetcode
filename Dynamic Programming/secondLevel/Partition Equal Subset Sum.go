package main

func canPartition(nums []int) bool {
	/*
	   	Input: nums = [1,5,11,5]
	      Output: true
	      Explanation: The array can be partitioned as [1, 5, 5] and [11].
	*/
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	dp := make([]int, sum+1)
	for _, num := range nums {
		for i := sum; i >= num; i-- {
			dp[i] = max(dp[i], dp[i-num]+num)
		}
	}
	if dp[sum] == sum {
		return true
	} else {
		return false
	}
}
