package main

func intersection(nums1 []int, nums2 []int) []int {
	retMap := map[int]bool{}
	m := map[int]bool{}
	for _, i := range nums1 {
		m[i] = true
	}
	for _, i := range nums2 {
		if m[i] == true {
			retMap[i] = true
		}
	}
	ret := []int{}
	for i := range retMap {
		ret = append(ret, i)
	}
	return ret
}
