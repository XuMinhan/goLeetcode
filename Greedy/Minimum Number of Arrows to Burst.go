package main

import "sort"

func main() {
	i := [][]int{
		{2, 8},
		{1, 6},
		{10, 12},
		{20, 22},
	}
	println(findMinArrowShots(i))
}

// [[10,16],[2,8],[1,6],[7,12]]
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] < points[j][0] {
			return true
		} else if points[i][0] > points[j][0] {
			return false
		} else {
			return points[i][1] < points[j][1]
		}
	})
	needNarrow := 1
	// 0 - 3
	//	 1 - 4
	//		4 - 6
	for index := 0; index < len(points); index++ {
		i := index + 1
		if i >= len(points) {
			return needNarrow
		}
		gang := points[index][1]
		for points[i][0] <= gang {
			if gang > points[i][1] {
				gang = points[i][1]
			}
			i++
			if i == len(points) {
				return needNarrow
			}
		}
		needNarrow++
		index = i - 1
	}
	return needNarrow
}
