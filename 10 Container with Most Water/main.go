package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(containerWithMostWaterOptimal(arr))

	arr = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(containerWithMostWaterOptimal(arr))

	arr = []int{3, 1, 2}
	fmt.Println(containerWithMostWaterOptimal(arr))

	arr = []int{0, 1}
	fmt.Println(containerWithMostWaterOptimal(arr))
}

func containerWithMostWaterBruteforce(arr []int) int {
	// using nested loops check with one element after other and store the maximum area
	// return the maximum area
	return 1
}

func containerWithMostWaterOptimal(height []int) int {
	left, right := 0, len(height)-1
	max := math.MinInt
	var area int

	for left < right {

		if height[left] < height[right] {
			area = height[left] * (right - left)
			max = Max(max, area)
			left++
		} else {
			area = height[right] * (right - left)
			max = Max(max, area)
			right--
		}

	}

	return max

}
func Max(max, val int) int {
	if max > val {
		return max
	}
	return val
}
