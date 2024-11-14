package main

import "fmt"

func main1() {
	strings := []string{"ll", "ll"}
	chars := commonChars(strings)
	fmt.Println(chars)
}
func commonChars(words []string) []string {
	m := map[rune]int{}
	first := true
	for _, word := range words {
		if first {
			for _, char := range word {
				m[char]++
			}
			first = false
		} else {
			newMap := map[rune]int{}
			for _, char := range word {
				if m[char] != 0 {
					newMap[char]++
					m[char]--
				}
			}
			m = newMap
		}
	}
	var strings []string
	for r, times := range m {
		for i := 0; i < times; i++ {
			strings = append(strings, string(r))
		}
	}
	return strings
}
