package main

import (
	"math"
)

func canCompleteCircuit(gas []int, cost []int) int {

	//  -1 3 -4 2
	do := make([]int, len(gas))
	for i := 0; i < len(gas); i++ {
		do[i] = gas[i] - cost[i]
	}
	total := 0
	for _, i := range do {
		total += i
	}
	if total < 0 {
		return -1
	}

	startIndex := 0
	sum := 0
	maxSum := math.MinInt
	retStartIndex := 0
	//9, -1, 2, -8, -1
	for i, num := range do {
		sum += num
		if sum < 0 {
			sum = 0
			startIndex = i + 1
			continue
		}
		if sum > maxSum {
			maxSum = sum
			retStartIndex = startIndex
		}
	}

	//  3 -1 -2 -3 4 -1

	endIndex := -1
	sum = 0
	minSum := math.MaxInt
	for i, num := range do {
		sum += num
		if sum > 0 {
			sum = 0
			continue
		}
		if sum < minSum {
			minSum = sum
			endIndex = i
		}
	}

	if maxSum+minSum > 0 {
		return retStartIndex
	} else {
		if endIndex == len(do)-1 {
			return 0
		} else {
			return endIndex + 1
		}
	}
}
