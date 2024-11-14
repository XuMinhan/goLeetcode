package main

import "fmt"

func main0() {
	bytes := []byte{'h', 'e'}
	reverseString(bytes)
	fmt.Println(bytes)
}
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
