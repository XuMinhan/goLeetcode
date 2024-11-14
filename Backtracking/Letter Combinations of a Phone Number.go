package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	m := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	var (
		path []byte
		ret  [][]byte
	)

	b := []byte(digits)
	var dfs func(digits []byte, start int)
	dfs = func(digits []byte, start int) {
		if start == len(digits) {
			ret = append(ret, append([]byte{}, path...))
			return
		}
		toBytes := m[digits[start]]
		for _, toByte := range toBytes {
			path = append(path, toByte)
			dfs(digits, start+1)
			path = path[:len(path)-1]
		}
	}
	dfs(b, 0)
	var retString []string
	for _, bytes := range ret {
		retString = append(retString, string(bytes))
	}
	return retString
}
