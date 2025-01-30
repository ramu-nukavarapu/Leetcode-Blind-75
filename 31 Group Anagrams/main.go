package main

import (
	"fmt"
	"strconv"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagramsOptimal(strs))

	strs = []string{""}
	fmt.Println((groupAnagramsOptimal(strs)))

	strs = []string{"a"}
	fmt.Println((groupAnagramsOptimal(strs)))

}

func groupAnagramsBruteforce() {
	// sort every element in the given array
	// add the elements that are same to list
	// add these lists to another list and return that list
}

func groupAnagramsOptimal(str []string) [][]string {
	group := make(map[string][]string)

	for _, val := range str {
		count := make([]int, 26)
		for _, value := range val {
			count[value-'a']++
		}

		key := ""

		for _, value := range count {
			key += strconv.Itoa(value) + "#"
		}

		group[key] = append(group[key], val)
	}

	result := make([][]string, 0)
	for _, val := range group {
		result = append(result, val)
	}

	return result
}
