package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(maximumSubarraySumOptimal(arr))

	arr = []int{2, 3, -8, 7, -1, 2, 3}
	fmt.Println(maximumSubarraySumOptimal(arr))

	arr = []int{-2, -4}
	fmt.Println(maximumSubarraySumOptimal(arr))

	arr = []int{5, 4, 1, 7, 8}
	fmt.Println(maximumSubarraySumOptimal(arr))

}

func maximumSubarraySumBruteforce(arr []int) int {
	// use nested loops start with first one and in the inner loop from that element iterate over and calculate maximum sum
	// check if the currentsum exceeds the maximum value if exceeds change the maximum value to the current sum
	return 0
}

func maximumSubarraySumOptimal(arr []int) int {
	max := math.MinInt
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		if sum > max {
			max = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	return max
}
