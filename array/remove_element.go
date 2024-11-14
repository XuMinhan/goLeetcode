package main

import "fmt"

func main1() {
	element := removeElement([]int{1, 2, 3, 1, 2}, 1)
	fmt.Println(element)
}

func removeElement(nums []int, val int) int {
	index := 0
	for _, num := range nums {
		if num != val {
			nums[index] = num
			index++
		}
	}
	return index
}
