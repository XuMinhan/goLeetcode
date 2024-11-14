package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var (
		path []int
		ret  [][]int
	)
	sort.Ints(nums)
	pre := -11
	var dfs func(nums []int, start int)
	dfs = func(nums []int, start int) {
		ret = append(ret, append([]int{}, path...))
		if start == len(nums) {
			return
		}
		for i := start; i < len(nums); i++ {
			if nums[i] == pre {
				continue
			}
			path = append(path, nums[i])
			dfs(nums, i+1)
			path = path[:len(path)-1]
			pre = nums[i]
		}

	}
	dfs(nums, 0)
	return ret
}
