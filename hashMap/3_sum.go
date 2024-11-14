package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{-1, 0, 1, 2, -1, -4}
	sum := threeSum(ints)
	fmt.Println(sum)
}
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	retMap := map[[3]int]bool{}
	ret := [][]int{}
	for firstIndex, num := range nums {
		target := -num
		secondIndex := firstIndex + 1
		thirdIndex := len(nums) - 1
		for secondIndex < thirdIndex {
			if nums[secondIndex]+nums[thirdIndex] < target {
				secondIndex++
			} else if nums[secondIndex]+nums[thirdIndex] > target {
				thirdIndex--
			} else {
				ints := [3]int{nums[firstIndex], nums[secondIndex], nums[thirdIndex]}
				_, exist := retMap[ints]
				if !exist {
					ret = append(ret, ints[:])
					retMap[ints] = true
				}
				secondIndex++
			}
		}
	}
	return ret
}
