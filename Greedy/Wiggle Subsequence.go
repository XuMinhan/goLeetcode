package main

func wiggleMaxLength(nums []int) int {
	up := 1
	down := 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 0 {
			up = max(up, down+1)
		} else if nums[i]-nums[i-1] < 0 {
			down = max(down, up+1)
		}
	}
	return max(up, down)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
