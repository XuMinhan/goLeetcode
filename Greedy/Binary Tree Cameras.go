package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	/*
		0 无监控，需要监控
		1 无监控，不需要监控
		2 有监控

		左右有一个 0 -> 2
		左右有一个 2 -> 1
		左右都是 1   -> 0
	*/
	ret := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 1
		}
		leftStatus := dfs(root.Left)
		rightStatus := dfs(root.Right)
		if leftStatus == 0 || rightStatus == 0 {
			ret++
			return 2
		}

		if leftStatus == 2 || rightStatus == 2 {
			return 1
		}
		return 0
	}
	if dfs(root) == 0 {
		ret++
	}
	if ret == 0 {
		return 1
	}
	return ret
}
