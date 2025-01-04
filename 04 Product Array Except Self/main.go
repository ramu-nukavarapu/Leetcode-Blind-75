package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(productArrayExceptSelfOptimal(arr))

	arr = []int{-1, 1, 0, -3, 3}
	fmt.Println(productArrayExceptSelfOptimal(arr))

	arr = []int{11, -2, 3, 7, 9, -10}
	fmt.Println(productArrayExceptSelfOptimal(arr))
}

func productArrayExceptSelfBruteforce(arr []int) []int {
	// using nested loops calculate the product except the element in current index
	// store the value into an output array and return it
	return []int{}
}

func productArrayExceptSelfOptimal(arr []int) []int {
	output := make([]int, len(arr))

	prefix := 1
	for i := 0; i < len(arr); i++ {
		output[i] = prefix
		prefix *= arr[i]
	}

	suffix := 1
	for i := len(arr) - 1; i >= 0; i-- {
		output[i] *= suffix
		suffix *= arr[i]
	}

	return output
}
