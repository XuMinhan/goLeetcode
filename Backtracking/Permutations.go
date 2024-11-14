package main

func permute(nums []int) [][]int {
	var (
		path []int
		res  [][]int
		spot map[int]bool
	)
	spot = make(map[int]bool)
	var dfs func(nums []int, start int)
	dfs = func(nums []int, start int) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if spot[i] == true {
				continue
			}
			spot[i] = true
			path = append(path, nums[i])
			dfs(nums, i+1)
			path = path[:len(path)-1]
			spot[i] = false
		}
	}
	dfs(nums, 0)
	return res
}
