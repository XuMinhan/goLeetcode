package main

func monotoneIncreasingDigits(n int) int {
	ints := []int{}
	sou := n
	// 1 1 0 1
	for n > 0 {
		ints = append([]int{n % 10}, ints...)
		n = n / 10
	}
	down := -1
	for i := range ints {
		if i == 0 {
			continue
		}
		if ints[i] < ints[i-1] {
			down = i
			break
		}
	}
	if down == -1 {
		return sou
	}
	// 2 0 -> 1 9

	index := down - 1
	ints[index] = ints[index] - 1
	for index > 0 && ints[index-1] == ints[index]+1 {
		ints[index] = 9
		index--
		ints[index] = ints[index] - 1
	}
	ret := 0
	for i, s := range ints {
		if i < down {
			ret = ret*10 + s
		} else {
			ret = ret*10 + 9
		}
	}
	return ret
}
