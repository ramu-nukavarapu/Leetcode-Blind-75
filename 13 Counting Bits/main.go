package main

import "fmt"

func main() {
	n := 0
	fmt.Println(countBitsOptimal(n))

	n = 2
	fmt.Println(countBitsOptimal(n))

	n = 5
	fmt.Println(countBitsOptimal(n))

	n = 10
	fmt.Println(countBitsOptimal(n))

	n = 102
	fmt.Println(countBitsOptimal(n))
}

func countingBitsBrutefoce(n int) []int {
	// iterate from 0 to n and apply the number of bits method to every number and store the output into an array
	// return the array
	return []int{}
}

func countBitsOptimal(n int) []int {
	dp := make([]int, n+1)
	offset := 1
	for i := 1; i <= n; i++ {
		if offset*2 == i {
			offset = i
		}
		dp[i] = 1 + dp[i-offset]
	}
	return dp
}
