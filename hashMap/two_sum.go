package main

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		i2, exist := m[target-num]
		if exist {
			return []int{i, i2}
		}
		m[num] = i
	}
	return []int{0}
}
