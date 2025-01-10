package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Printf("Sum := %d\n", getSum(a, b))

	a, b = 7, 4
	fmt.Printf("Sum := %d\n", getSum(a, b))

	a, b = 19, 23
	fmt.Printf("Sum := %d\n", getSum(a, b))

}

func getSum(a int, b int) int {
	for b != 0 {
		temp := (a & b) << 1
		a = a ^ b
		b = temp
	}
	return a
}
