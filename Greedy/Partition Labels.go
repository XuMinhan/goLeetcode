package main

// "ababcbacadefegdehijhklij"
func partitionLabels(s string) (partition []int) {
	lastPos := [26]int{}
	for i, char := range s {
		lastPos[char-'a'] = i
	}
	start, end := 0, 0
	for i, char := range s {
		if lastPos[char-'a'] > end {
			end = lastPos[char-'a']
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = i + 1
		}
	}
	return partition
}
