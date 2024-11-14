package main

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g) // 1 2 3 4 6
	sort.Ints(s) // 2 3 4 6 10
	len1 := len(g)
	len2 := len(s)
	i2 := 0
	ret := 0
	for i := 0; i < len1; i++ {
		if i2 == len(s) {
			return i
		}
		for j := i2; j < len2; j++ {
			if s[j] >= g[i] {
				i2 = j + 1
				ret++
				break
			}
		}
	}
	println(111)
	return ret
}
