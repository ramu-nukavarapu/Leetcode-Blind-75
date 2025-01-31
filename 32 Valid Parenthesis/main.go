package main

import "fmt"

func main() {
	s := "()"
	fmt.Println(isValid(s))

	s = "()[]{}"
	fmt.Println(isValid(s))

	s = "(]"
	fmt.Println(isValid(s))

	s = "([])"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
