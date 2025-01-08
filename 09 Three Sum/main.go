package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-2, -1, 0, 1}
	fmt.Println(threeSumOptimal(arr))

	arr = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSumOptimal(arr))

	arr = []int{0, 1, 1}
	fmt.Println(threeSumOptimal(arr))

	arr = []int{0, 0, 0}
	fmt.Println(threeSumOptimal(arr))
}

func threeSumBruteforce(arr []int) [][]int {
	// using nested loops,check all possible triplets and select the unique triplets
	return [][]int{}
}

func threeSumOptimal(arr []int) [][]int {
	sort.Ints(arr)
	list := make([][]int, 0)

	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		j := i + 1
		k := len(arr) - 1

		for j < k {
			sum := arr[i] + arr[j] + arr[k]
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				sol := []int{arr[i], arr[j], arr[k]}
				list = append(list, sol)
				j++
				k--
				for j < k && arr[j] == arr[j+1] {
					j++
				}
				for j < k && arr[k] == arr[k-1] {
					k--
				}
			}
		}
	}
	return list
}
