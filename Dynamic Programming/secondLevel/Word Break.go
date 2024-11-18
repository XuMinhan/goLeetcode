package main

func main() {
	println(wordBreak("lec", []string{"le", "c"}))
}
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	// dp[i] = dp[i-len(s)] && wordDictSet[s]
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			s1 := s[j:i]
			dp[i] = dp[i-len(s1)] && wordDictSet[s1]
			if dp[i] {
				break
			}
		}
	}
	return dp[len(s)]
}
