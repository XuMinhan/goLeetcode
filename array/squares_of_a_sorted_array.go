package main

func sortedSquares(nums []int) []int {
	bottom := 0
	for i, num := range nums {
		if num > 0 {
			bottom = i
			break
		}
	}
	if nums[len(nums)-1] <= 0 {
		bottom = len(nums) - 1
		var ret []int
		index := 0
		for bottom > -1 {
			ret[index] = nums[bottom] * nums[bottom]
			bottom--
			index++
		}
		return ret
	}
	left := bottom - 1
	var ret []int
	index := 0
	for bottom <= len(nums)-1 && left >= 0 {
		if -nums[left] < nums[bottom] {
			ret[index] = -nums[left] * -nums[left]
			left--
		} else {
			ret[index] = nums[bottom] * nums[bottom]
			bottom++
		}
		index++
	}
	if left == -1 {
		for bottom < len(nums) {
			ret[index] = -nums[bottom] * -nums[bottom]
			bottom++
			index++
		}
	} else {
		for left > -1 {
			ret[index] = -nums[left] * -nums[left]
			left--
			index++
		}
	}
	return ret
}
