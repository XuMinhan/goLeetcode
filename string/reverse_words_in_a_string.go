package main

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	split := strings.Fields(s)
	l := len(split)
	for i := 0; i < l/2; i++ {
		split[i], split[l-i-1] = split[l-i-1], split[i]
	}
	return strings.Join(split, " ")
}
