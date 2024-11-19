package main

func trap(height []int) int {
	/*
	   	Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
	      Output: 6
	*/
	leftMaxHigh := make([]int, len(height))
	rightMaxHigh := make([]int, len(height))
	max0 := 0
	for i := range height {
		if i == 0 {
			continue
		}
		max0 = max(max0, height[i-1])
		leftMaxHigh[i] = max0
	}
	max0 = 0
	for i := len(height) - 2; i >= 0; i-- {
		max0 = max(max0, height[i+1])
		rightMaxHigh[i] = max0
	}
	ret := 0
	for i := 0; i < len(height); i++ {
		tmp := min(leftMaxHigh[i], rightMaxHigh[i]) - height[i]
		if tmp > 0 {
			ret += tmp
		}
	}
	return ret
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
