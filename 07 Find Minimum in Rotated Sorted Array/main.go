package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMinimuminRotatedSortArrayOptimal(nums))

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMinimuminRotatedSortArrayOptimal(nums))

	nums = []int{2, 1}
	fmt.Println(findMinimuminRotatedSortArrayOptimal(nums))

	nums = []int{11, 13, 15, 17}
	fmt.Println(findMinimuminRotatedSortArrayOptimal(nums))

}

func findMinimuminRotatedSortArrayBruteforce(arr []int) int {
	// go through the array, and check one after if the current element is lesser than the minimum element
	// update minimum with current one repeat the process till the array ends
	// return the minimum
	return 1
}

func findMinimuminRotatedSortArrayOptimal(nums []int) int {
	left, right := 0, len(nums)-1
	var min int = math.MaxInt
	for left <= right {
		mid := (left + right) / 2
		if nums[left] <= nums[mid] {
			min = checkMin(min, nums[left])
			left = mid + 1

		} else {
			min = checkMin(min, nums[mid])
			right = mid - 1
		}
	}
	return min
}
func checkMin(min, new int) int {
	if min < new {
		return min
	}
	return new
}
