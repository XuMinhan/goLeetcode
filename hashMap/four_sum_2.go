package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m1 := map[int]int{}
	for _, i := range nums1 {
		for _, i2 := range nums2 {
			m1[i+i2]++
		}
	}

	m2 := map[int]int{}
	for _, i := range nums3 {
		for _, i2 := range nums4 {
			m2[i+i2]++
		}
	}
	ret := 0
	for sum1, times := range m1 {
		ret += m2[-sum1] * times
	}

	return ret
}
