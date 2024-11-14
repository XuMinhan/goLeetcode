package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	ret := [][]int{}
	retMap := map[[4]int]bool{}
	for i1, num1 := range nums[:len(nums)-3] {
		for i20, num2 := range nums[i1+1 : len(nums)-2] {
			i2 := i20 + i1 + 1
			target34 := target - num2 - num1
			index3 := i2 + 1
			index4 := len(nums) - 1
			for index3 < index4 {
				if nums[index3]+nums[index4] < target34 {
					index3++
				} else if nums[index3]+nums[index4] > target34 {
					index4--
				} else {
					ints := [4]int{nums[i1], nums[i2], nums[index3], nums[index4]}
					_, exist := retMap[ints]
					if !exist {
						ret = append(ret, ints[:])
						retMap[ints] = true
					}
					index4--
					index3++

				}
			}
		}
	}
	return ret
}
