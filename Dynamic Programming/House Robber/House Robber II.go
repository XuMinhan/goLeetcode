package main

func main() {
	println(rob([]int{1, 2, 1, 0}))
}

// 1 2 1 0

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	/*
				2,3,2
		3 false true
		3 true false

		a bcdefgh i

		如果 a 固定不偷，后面的随意，最大为 m
			b和i都不偷，代表 c和h都偷
				加上 a 就是答案
			其他情况
				作为 ret0

		如果 i 固定不偷，前面的随意，最大为 n
			a和h都不偷，代表 b和g都偷
				加上 i 就是答案
			其他情况
				作为 ret1
	*/
	max0, headStolen0, tailStolen0 := robRob(nums[:len(nums)-1])
	max1, headStolen1, tailStolen1 := robRob(nums[1:])
	if !headStolen0 && !tailStolen0 {
		return max(max0+nums[0], max1)
	}
	if !headStolen1 && !tailStolen1 {
		return max(max1+nums[len(nums)-1], max0)
	}
	return max(max1, max0)
}
func robRob(nums []int) (max int, headStolen, tailStolen bool) {
	dp := make([][][]int, len(nums))
	/*


		[2,1,1,2]

		[2,1] [0,0]
		[1,0] [2,1]
		[3,1] [2,1]
		[4,1] [3,1]
	*/
	for i := range dp {
		dp[i] = make([][]int, 2)
		for i2 := range dp[i] {
			dp[i][i2] = make([]int, 2)
		}
	}
	dp[0][0][0] = nums[0]
	dp[0][0][1] = 1
	dp[0][1][0] = 0
	dp[0][1][1] = 0
	for i := 1; i < len(nums); i++ {
		dp[i][0][0] = dp[i-1][1][0] + nums[i]
		dp[i][0][1] = dp[i-1][1][1]
		if dp[i-1][0][0] >= dp[i-1][1][0] {
			dp[i][1][0] = dp[i-1][0][0]
			dp[i][1][1] = dp[i-1][0][1]
		} else {
			dp[i][1][0] = dp[i-1][1][0]
			dp[i][1][1] = dp[i-1][1][1]
		}
	}
	if dp[len(nums)-1][0][0] >= dp[len(nums)-1][1][0] {
		return dp[len(nums)-1][0][0], dp[len(nums)-1][0][1] == 1, true
	} else {
		return dp[len(nums)-1][1][0], dp[len(nums)-1][1][1] == 1, false
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
