package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "bbbbb"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	maxLength := 0
	lastIndex := make([]int, 128)

	start := 0
	for end := 0; end < n; end++ {
		currentChar := s[end]
		if lastIndex[currentChar] > start {
			start = lastIndex[currentChar]
		}
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
		lastIndex[currentChar] = end + 1
	}

	return maxLength
}
