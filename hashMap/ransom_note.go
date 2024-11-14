package main

func canConstruct(ransomNote string, magazine string) bool {
	m := map[rune]int{}
	for _, i := range magazine {
		m[i]++
	}
	for _, i := range ransomNote {
		if m[i] == 0 {
			return false
		}
		m[i]--
	}
	return true
}
