package main

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	/*
		people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
		4	5	6	7
		1	2	1	2
		_ _ _ _ 4 _ _
	*/
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] < people[j][0] {
			return true
		} else if people[i][0] > people[j][0] {
			return false
		} else {
			return people[i][1] > people[j][1]
		}
	})
	var insert func(nums [][]int, pos int, num []int)
	insert = func(nums [][]int, pos int, numForInsert []int) {
		index := 0
		for i, num := range nums {
			if num != nil {
				continue
			}
			if index == pos {
				nums[i] = numForInsert
			}
			index++
		}
	}
	retPeople := make([][]int, len(people))
	for _, person := range people {
		insert(retPeople, person[1], person)
	}
	return retPeople
}
