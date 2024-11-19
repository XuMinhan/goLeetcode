package Monotonic_Stack

import "container/list"

func dailyTemperatures(temperatures []int) []int {
	/*
		输入: temperatures = [73,74,75,71,69,72,76,73]
		输出: [1,1,4,2,1,1,0,0]
	*/
	stack := list.New()
	ret := make([]int, len(temperatures))
	for i, temperature := range temperatures {
		for stack.Len() != 0 && temperature > temperatures[stack.Front().Value.(int)] {
			out := stack.Front().Value.(int)
			ret[out] = i - out
			stack.Remove(stack.Front())
		}
		stack.PushFront(i)
	}
	return ret
}
