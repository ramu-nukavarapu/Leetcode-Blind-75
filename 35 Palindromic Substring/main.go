package main

import "fmt"

func main() {
	s := "abc"
	fmt.Println(countSubstrings(s))

	s = "aaa"
	fmt.Println(countSubstrings(s))
}

func countSubstrings(s string) int {
	n := len(s)
	res := 0

	for i := 0; i < n; i++ {
		start, end := i, i
		for start >= 0 && end < n && s[start] == s[end] {
			res++
			start--
			end++
		}

		start, end = i-1, i
		for start >= 0 && end < n && s[start] == s[end] {
			res++
			start--
			end++
		}
	}

	return res
}
