package main

import "math"

func maxSubArray(nums []int) int {
	sum := 0
	ret := 0
	dayu0 := false
	for _, num := range nums {
		sum += num
		if sum < 0 {
			sum = 0
			continue
		}
		if sum > ret {
			dayu0 = true
			ret = sum
		}
	}
	if dayu0 == true {
		return ret
	}
	ret = math.MinInt
	for _, num := range nums {
		if num > ret {
			ret = num
		}
	}
	return ret
}
