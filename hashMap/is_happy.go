package main

func isHappy(n int) bool {
	m := map[int]bool{}
	for m[n] == false {
		m[n] = true
		n = step(n)
	}
	return n == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
