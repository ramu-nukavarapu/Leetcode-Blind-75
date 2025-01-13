package main

import "fmt"

func main() {
	arr := []int{3, 0, 1}
	fmt.Println(missingNumberOptimal(arr))

	arr = []int{1}
	fmt.Println(missingNumberOptimal(arr))

	arr = []int{0, 1}
	fmt.Println(missingNumberOptimal(arr))

	arr = []int{2, 1}
	fmt.Println(missingNumberOptimal(arr))

	arr = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumberOptimal(arr))

}

func missingNumberBruteforce(arr []int) int {
	// sort the array and difference between val and index
	// if the difference is greater than 0 return the index
	return 0
}

func missingNumberOptimal(nums []int) int {
	val := 0

	for i := 0; i < len(nums); i++ {
		val = val ^ nums[i] ^ i
	}
	return val ^ len(nums)
}
