package main

import "fmt"

func main() {
	n := 128
	fmt.Println(numberof1BitsOptimal(n))

	n = 12
	fmt.Println(numberof1BitsOptimal(n))

	n = 0
	fmt.Println(numberof1BitsOptimal(n))

	n = 128233421
	fmt.Println(numberof1BitsOptimal(n))

}

func numberof1BitsBruteforce(n int) int {
	// convert the integer into binary string
	// iterate over it and increment counter by 1 when the value at current index is 1
	// return the counter
	return 0
}

func numberof1BitsOptimal(n int) int {
	res := 0
	for n > 0 {
		n = n & (n - 1)
		res++
	}
	return res
}
