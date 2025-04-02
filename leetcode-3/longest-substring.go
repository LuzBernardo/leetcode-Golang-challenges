package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	charIndex := make(map[byte]int)
	maxLen := 0
	start := 0

	for i := 0; i < len(s); i++ {
		if lastIndex, ok := charIndex[s[i]]; ok && lastIndex >= start {
			start = lastIndex + 1
		}
		charIndex[s[i]] = i
		currLen := i - start + 1
		if currLen > maxLen {
			maxLen = currLen
		}
	}
	return maxLen
}
