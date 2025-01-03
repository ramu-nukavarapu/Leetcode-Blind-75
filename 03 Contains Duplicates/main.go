package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicatesOptimal(nums))

	nums = []int{1, 2, 3, 4}
	fmt.Println(containsDuplicatesOptimal(nums))

	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicatesOptimal(nums))
}

func containsDuplicatesBruteforce(nums []int) bool {
	// sort the array. check adjacent elements difference
	// if difference == 0 then return false else return true
	return true
}

func containsDuplicatesOptimal(nums []int) bool {
	hashmap := make(map[int]struct{}, len(nums))

	for i := 0; i < len(nums); i++ {
		if _, ok := hashmap[nums[i]]; ok {
			return true
		}
		hashmap[nums[i]] = struct{}{}
	}
	return false
}
