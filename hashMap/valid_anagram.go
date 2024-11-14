package main

func main0() {
	println(isAnagram("ab", "a"))
}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[rune]int{}
	for _, i := range s {
		m[i]++
	}
	for _, i := range t {
		if m[i] == 0 {
			return false
		} else {
			m[i]--
		}
	}
	return true
}
