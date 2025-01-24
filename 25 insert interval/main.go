package main

import "fmt"

func main() {
	arr := [][]int{{1, 3}, {6, 9}}
	new := []int{2, 5}
	fmt.Println(insert(arr, new))

	arr = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	new = []int{4, 8}
	fmt.Println(insert(arr, new))

	arr = [][]int{{2, 3}, {6, 9}}
	new = []int{0, 1}
	fmt.Println(insert(arr, new))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)

	for element := range intervals {
		if intervals[element][0] > newInterval[1] {
			result = append(result, newInterval)
			result = append(result, intervals[element:]...)
			return result
		} else if intervals[element][1] < newInterval[0] {
			result = append(result, intervals[element])
		} else {
			newInterval[0] = min(intervals[element][0], newInterval[0])
			newInterval[1] = max(intervals[element][1], newInterval[1])
		}
	}
	result = append(result, newInterval)
	return result
}
