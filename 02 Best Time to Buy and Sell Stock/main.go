package main

import "fmt"

func main() {
	prices := []int{10, 3, 4, 5, 2, 9, 12}
	fmt.Println(maxProfitOptimal(prices))

	prices = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfitOptimal(prices))

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfitOptimal(prices))
}

func maxProfitBruteforce(prices []int) int {
	// using nested loops, select one element and check with other elements
	// if the difference b/w them is positive check with currentprofit if it is greater then set currentprofit = difference
	return 0
}

func maxProfitOptimal(prices []int) int {
	profit := 0

	first, second := 0, 1
	current := 0

	for second < len(prices) {
		if prices[first] < prices[second] {
			current = prices[second] - prices[first]
			second++
			if current > profit {
				profit = current
			}
		} else {
			first = second
			second++
		}
	}

	return profit
}
