package main

import (
	"fmt"
	"unicode"
)

func main() {

	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))

	s = "race a car"
	fmt.Println(isPalindrome(s))

	s = " "
	fmt.Println(isPalindrome(s))

}

func isPalindrome(s string) bool {
	start, end := 0, len(s)-1

	for start < end {
		if !(unicode.IsDigit(rune(s[start])) || unicode.IsLetter(rune(s[start]))) {
			start++
			continue
		} else if !(unicode.IsDigit(rune(s[end])) || unicode.IsLetter(rune(s[end]))) {
			end--
			continue
		} else if unicode.ToLower(rune(s[start])) != unicode.ToLower(rune(s[end])) {
			return false
		} else {
			start++
			end--
		}
	}
	return true
}
