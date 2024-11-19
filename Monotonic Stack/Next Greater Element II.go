package Monotonic_Stack

import "container/list"

func main() {
	nextGreaterElements([]int{1, 2, 1})
}
func nextGreaterElements(nums []int) []int {

	sLen := len(nums)
	nums = append(nums, nums...)
	stack := list.New()
	ret := make([]int, len(nums))
	for i := range ret {
		if ret[i] == 0 {
			ret[i] = -1
		}
	}
	for i, temperature := range nums {
		for stack.Len() != 0 && temperature > nums[stack.Front().Value.(int)] {
			out := stack.Front().Value.(int)
			ret[out] = nums[i]
			stack.Remove(stack.Front())
		}
		stack.PushFront(i)
	}

	return ret[:sLen]
}
