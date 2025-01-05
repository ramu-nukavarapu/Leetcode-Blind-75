package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 3, 4}
	fmt.Println(maximumSubarrayProductOptimal(arr))

	arr = []int{0}
	fmt.Println(maximumSubarrayProductOptimal(arr))

	arr = []int{-1, 3, 2, 1, -5}
	fmt.Println(maximumSubarrayProductOptimal(arr))

	arr = []int{2, 3, -4}
	fmt.Println(maximumSubarrayProductOptimal(arr))
}

func maximumSubarrayProductBruteforce(nums []int) int {
	// using nested loops check all the sub arrays product and return the maximum one
	return 1
}

func maximumSubarrayProductOptimal(nums []int) int {
	max := math.MinInt

	prefix := 1
	suffix := 1
	i := 0
	for i < len(nums) {
		prefix *= nums[i]
		suffix *= nums[len(nums)-i-1]

		if prefix > max {
			max = prefix
		}

		if suffix > max {
			max = suffix
		}

		if prefix == 0 {
			prefix = 1
		}
		if suffix == 0 {
			suffix = 1
		}
		i++
	}
	return max
}
