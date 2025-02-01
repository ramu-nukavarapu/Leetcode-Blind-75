package main

import "fmt"

func main() {
	s := "ABAB"
	k := 2
	fmt.Println(characterReplacement(s, k))

	s = "AABABBA"
	k = 1
	fmt.Println(characterReplacement(s, k))
}

func characterReplacement(s string, k int) int {
	freqs := make(map[byte]int)
	res, i, maxFreq := 0, 0, 0

	for j := 0; j < len(s); j++ {
		freqs[s[j]]++
		if freqs[s[j]] > maxFreq {
			maxFreq = freqs[s[j]]
		}

		for (j-i+1)-maxFreq > k {
			freqs[s[i]]--
			i++
		}

		if j-i+1 > res {
			res = j - i + 1
		}
	}

	return res
}
