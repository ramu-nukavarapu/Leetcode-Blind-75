package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))

	s = "rat"
	t = "car"
	fmt.Println(isAnagram(s, t))
}

func isAnagramBruteforce() {
	// sort both the strings
	// check the strings are equal or not
	//if equal return true else return false
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	alphabets := make([]int, 26)

	for i := 0; i < len(s); i++ {
		alphabets[s[i]-'a']++
		alphabets[t[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if alphabets[i] != 0 {
			return false
		}
	}
	return true
}
