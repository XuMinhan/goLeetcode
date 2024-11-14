package main

func combinationSum(candidates []int, target int) [][]int {
	var (
		path []int
		res  [][]int
		sum  int
	)
	var dfs func(candidates []int, target int, start int, sum int)
	dfs = func(candidates []int, target int, start int, sum int) {
		if target < sum {
			return
		}
		if target == sum {
			res = append(res, append([]int{}, path...))
			return
		}
		for i, candidate := range candidates {
			if i < start {
				continue
			}
			sum += candidate
			path = append(path, candidate)
			dfs(candidates, target, i, sum)
			path = path[:len(path)-1]
			sum -= candidate
		}
	}
	dfs(candidates, target, 0, sum)
	return res
}
