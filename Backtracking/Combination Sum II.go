package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var (
		path []int
		res  [][]int
		sum  int
	)
	sort.Ints(candidates)
	var dfs func(candidates []int, target int, start int, sum int)
	dfs = func(candidates []int, target int, start int, sum int) {
		if target < sum {
			return
		}
		if target == sum {
			res = append(res, append([]int{}, path...))
			return
		}
		var pre int
		for i, candidate := range candidates {
			if candidate == pre {
				continue
			}
			if i < start {
				continue
			}
			sum += candidate
			path = append(path, candidate)
			dfs(candidates, target, i+1, sum)
			path = path[:len(path)-1]
			sum -= candidate
			pre = candidate
		}
	}
	dfs(candidates, target, 0, sum)
	return res
}
