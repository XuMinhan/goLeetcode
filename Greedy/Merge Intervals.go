package main

import (
	"fmt"
	"sort"
)

func main0() {
	/*
		[1,3],[2,6],[8,10],[15,18]
	*/
	i := merge([][]int{
		{1, 4},
		{1, 4},
	})
	fmt.Println(i)

}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] > intervals[j][1]
		}
	})
	ret := [][]int{}
	for i := 0; i < len(intervals); i++ {
		if i == len(intervals)-1 {
			ret = append(ret, []int{intervals[i][0], intervals[i][1]})
			return ret
		}
		gang := intervals[i][1]
		index := i + 1
		for gang >= intervals[index][0] {
			if gang < intervals[index][1] {
				gang = intervals[index][1]
			}
			index++
			if index == len(intervals) {
				ret = append(ret, []int{intervals[i][0], gang})
				return ret
			}
		}
		ret = append(ret, []int{intervals[i][0], gang})
		i = index - 1
	}
	return ret
}
