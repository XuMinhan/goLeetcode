package main

func combine(n int, k int) [][]int {
	var (
		path []int
		res  [][]int
	)
	var dfs func(n int, k int, start int)
	dfs = func(n int, k int, start int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			dfs(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(n, k, 1)
	return res
}
