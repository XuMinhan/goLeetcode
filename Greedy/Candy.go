package main

func candy(ratings []int) (ans int) {
	candies := make([]int, len(ratings))
	for i, rating := range ratings {
		if i == 0 {
			continue
		}
		if rating > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candies[i+1]+1 > candies[i] {
				candies[i] = candies[i+1] + 1
			}
		}
	}
	ret := 0
	for _, i := range candies {
		ret += i
	}
	return ret + len(candies)

}
