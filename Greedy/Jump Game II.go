package main

/*
Example 1:

Input: nums = [2,3,1,1,4]
Input: nums = [2,3,1]
Output: 2
*/
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	maxReach := 0
	chosenPos := 0
	tmp := 0
	ret := 1
	for {
		chosenPos = tmp
		for i := 1; i <= nums[chosenPos]; i++ {
			pos := i + chosenPos
			if pos >= len(nums)-1 {
				return ret
			}
			if pos+nums[pos] > maxReach {
				maxReach = pos + nums[pos]
				tmp = pos
			}
		}
		ret++
	}
}
