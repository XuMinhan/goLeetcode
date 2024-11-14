package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	/*
		Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
		Output: 1
		Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.
	*/
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return intervals[i][0] > intervals[j][0]
		}
	})
	gang := intervals[0][1]
	ret := 0
	for i, interval := range intervals {
		if i == 0 {
			continue
		}
		if interval[0] < gang {
			ret++
		} else {
			gang = interval[1]
		}
	}
	return ret
}
