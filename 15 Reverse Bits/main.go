package main

import "fmt"

func main() {
	var num uint32

	num = 3
	fmt.Println(reverseBitsOptimal(num))

	num = 7
	fmt.Println(reverseBitsOptimal(num))

	num = 2147483648
	fmt.Println(reverseBitsOptimal(num))

	num = 11
	fmt.Println(reverseBitsOptimal(num))
}

func reverseBitsBruteforce(num uint32) uint32 {
	// convert the number into binary string
	// reverse the string
	// convert the binary string into number and return it
	return 0
}

func reverseBitsOptimal(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res<<1 | (num & 1)
		num = num >> 1
	}
	return res
}
