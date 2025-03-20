package main

func wordBreak(s string, wordDict []string) bool {
	memo := make(map[string]bool)

	var dp func(string) bool
	dp = func(s string) bool {
		if len(s) == 0 {
			return true
		}
		if val, ok := memo[s]; ok {
			return val
		}

		for _, word := range wordDict {
			if len(s) >= len(word) && s[0:len(word)] == word &&
				dp(s[len(word):]) {
				memo[s] = true
				return true
			}
		}

		memo[s] = false
		return false
	}

	return dp(s)
}
