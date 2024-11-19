package main

import "container/list"

func main() {
	println(largestRectangleArea([]int{1}))
}
func largestRectangleArea(heights []int) int {
	l0 := list.New()
	rightFirstSmaller := make([]int, len(heights))
	for i := range rightFirstSmaller {
		rightFirstSmaller[i] = -1
	}
	leftFirstSmaller := make([]int, len(heights))
	for i := range leftFirstSmaller {
		leftFirstSmaller[i] = -1
	}
	for i, height := range heights {
		for l0.Len() != 0 && height < heights[l0.Front().Value.(int)] {
			popIndex := l0.Front().Value.(int)
			rightFirstSmaller[popIndex] = i
			l0.Remove(l0.Front())
		}
		l0.PushFront(i)
	}
	rightFirstSmaller[len(heights)-1] = len(heights)
	leftFirstSmaller[0] = -1
	for i := len(heights) - 1; i >= 0; i-- {
		for l0.Len() != 0 && heights[i] < heights[l0.Front().Value.(int)] {
			popIndex := l0.Front().Value.(int)
			leftFirstSmaller[popIndex] = i
			l0.Remove(l0.Front())
		}
		l0.PushFront(i)
	}
	for i, i2 := range rightFirstSmaller {
		if i2 == -1 {
			rightFirstSmaller[i] = len(heights)
		}
	}

	ret := 0
	for i := 0; i < len(heights); i++ {
		ret = max(ret, heights[i]*(rightFirstSmaller[i]-leftFirstSmaller[i]-1))
	}
	return ret
}
