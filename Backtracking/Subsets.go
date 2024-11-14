package main

func subsets(nums []int) [][]int {
	var (
		path []int
		ret  [][]int
	)
	var dfs func(nums []int, start int)
	dfs = func(nums []int, start int) {
		ret = append(ret, append([]int{}, path...))
		if start == len(nums) {
			return
		}
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(nums, 0)
	return ret
}
