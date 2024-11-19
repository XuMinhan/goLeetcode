package main

import "container/list"

func nextGreaterElement(nums1, nums2 []int) []int {
	m := map[int]int{}
	stack := list.New()
	for _, num := range nums2 {
		for stack.Len() != 0 && num > stack.Front().Value.(int) {
			m[stack.Front().Value.(int)] = num
			stack.Remove(stack.Front())
		}
		stack.PushFront(num)
	}
	ret := make([]int, len(nums1))
	for i, i2 := range nums1 {
		val, exist := m[i2]
		if exist {
			ret[i] = val
		} else {
			ret[i] = -1
		}
	}
	return ret
}
