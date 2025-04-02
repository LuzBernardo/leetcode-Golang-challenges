package main

import "fmt"

func longestSubstring(s string, k int) int {
	var findLongest func(start, end int) int
	findLongest = func(start, end int) int {
		if start >= end {
			return 0
		}
		if end-start < k {
			return 0
		}

		freq := make(map[byte]int)
		for i := start; i < end; i++ {
			freq[s[i]]++
		}

		valid := true
		for _, count := range freq {
			if count < k {
				valid = false
				break
			}
		}
		if valid {
			return end - start
		}

		maxLen := 0
		lastStart := start
		for i := start; i < end; i++ {
			if freq[s[i]] < k {
				currLen := findLongest(lastStart, i)
				if currLen > maxLen {
					maxLen = currLen
				}
				lastStart = i + 1
			}
		}

		if lastStart < end {
			currLen := findLongest(lastStart, end)
			if currLen > maxLen {
				maxLen = currLen
			}
		}

		return maxLen
	}

	return findLongest(0, len(s))
}

func main() {
	s := "aaaab"

	fmt.Println(longestSubstring(s, 3))
}
