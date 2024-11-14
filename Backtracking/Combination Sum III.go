package main

func combinationSum3(k int, n int) [][]int {
	var (
		path []int
		res  [][]int
		sum  int
	)
	var dfs func(k int, n int, start int)
	dfs = func(k int, n int, start int) {
		if len(path) == k {
			if n == sum {
				res = append(res, append([]int{}, path...))
			}
			return
		}
		for i := start; i <= 9; i++ {
			path = append(path, i)
			sum += i
			dfs(k, n, i+1)
			sum -= i
			path = path[:len(path)-1]
		}
	}
	dfs(k, n, 1)
	return res
}
