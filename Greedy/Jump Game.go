package main

func canJump(nums []int) bool {
	var jump func(nums []int, start int, unlock int) (maxReach int)
	jump = func(nums []int, start int, unlock int) (maxReach int) {
		for i := start; i <= unlock; i++ {
			if nums[i]+i > maxReach {
				maxReach = nums[i] + i
			}
		}
		return maxReach
	}
	maxReach := 0
	pre := 0
	for maxReach < len(nums)-1 {
		unlock := maxReach
		maxReach = jump(nums, pre, unlock)
		pre = unlock + 1
		if maxReach == unlock {
			return false
		}
	}
	return true
}
