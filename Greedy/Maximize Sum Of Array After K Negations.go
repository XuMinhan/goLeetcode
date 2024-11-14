package main

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	/*
		   	Example 1:

		      Input: nums = [4,2,3], k = 1
		      Output: 5
		      Explanation: Choose index 1 and nums becomes [4,-2,3].

		-9 -8 -7 -6  1 2 3 4 5
	*/
	sort.Ints(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	zeroIndex := -1
	for i, num := range nums {
		if num < 0 {
			zeroIndex = i
		}
	}
	if zeroIndex == -1 {
		if k%2 == 0 {
			return sum
		} else {
			return sum - 2*nums[0]
		}
	}
	if zeroIndex+1 >= k {
		for i := 0; i < k; i++ {
			sum -= 2 * nums[i]
		}
		return sum
	}
	for i := 0; i < zeroIndex+1; i++ {
		sum -= 2 * nums[i]
	}
	if (k-zeroIndex-1)%2 == 0 {
		return sum
	}
	if zeroIndex == len(nums)-1 {
		return sum + 2*nums[zeroIndex]
	}
	// -1 0 2 3。3
	if nums[zeroIndex]+nums[zeroIndex+1] > 0 {
		return sum + 2*nums[zeroIndex]
	} else {
		return sum - 2*nums[zeroIndex+1]
	}
}
