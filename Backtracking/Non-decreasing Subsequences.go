package main

func findSubsequences(nums []int) [][]int {
	var (
		path []int
		ret  [][]int
	)

	var dfs func(nums []int, start int)
	dfs = func(nums []int, start int) {
		if len(path) > 1 {
			if path[len(path)-1] < path[len(path)-2] {
				return
			}
			ret = append(ret, append([]int{}, path...))
		}
		if start == len(nums) {
			return
		}
		m := map[int]bool{}
		for i := start; i < len(nums); i++ {
			if m[nums[i]] == true {
				continue
			}
			m[nums[i]] = true
			path = append(path, nums[i])
			dfs(nums, i+1)
			path = path[:len(path)-1]
		}

	}
	dfs(nums, 0)
	return ret
}
