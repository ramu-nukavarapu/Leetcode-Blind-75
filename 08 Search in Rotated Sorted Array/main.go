package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 0, 1, 2}
	target := 0
	fmt.Println(searchinRotatedSortedArrayOptimal(nums, target))

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 1
	fmt.Println(searchinRotatedSortedArrayOptimal(nums, target))

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 3
	fmt.Println(searchinRotatedSortedArrayOptimal(nums, target))

}

func searchinRotatedSortedArrayBruteforce(nums []int, target int) int {
	// check one after other element with target if target matches with current element return current element
	// else return -1
	return -1
}

func searchinRotatedSortedArrayOptimal(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
