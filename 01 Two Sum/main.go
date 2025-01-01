package main

import "fmt"

func main() {
	arr := []int{2, 3, 11, 15, 7}
	target := 9
	fmt.Println(twoSumOptimal(arr, target))
}

func twoSumBruteforce(arr []int, target int) []int {
	sol := []int{-1, -1}
	// use nested loops select one element and re-iterate the remaining elements
	// if there sum equals to target return their indexes else return [-1,-1] as not found
	return sol
}

func twoSumOptimal(arr []int, target int) []int {

	if len(arr) == 1 && arr[0]+arr[0] == target {
		return []int{0, 0}
	}

	needed := make(map[int]int, 0)

	need := -1
	for i := 0; i < len(arr); i++ {
		need = target - arr[i]
		val, ok := needed[need]
		if ok == true {
			return []int{i, val}
		}
		needed[arr[i]] = i
	}

	return []int{-1, -1}
}
