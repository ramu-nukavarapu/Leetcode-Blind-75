package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{1, 3}, {6, 9}, {2, 5}}
	fmt.Println(merge(arr))

	arr = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}, {4, 8}}
	fmt.Println(merge(arr))

	arr = [][]int{{2, 3}, {6, 9}, {0, 1}}
	fmt.Println(merge(arr))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	output := make([][]int, 1)
	output[0] = intervals[0]

	for _, val := range intervals[1:] {
		if output[len(output)-1][1] >= val[0] {
			output[len(output)-1][1] = max(output[len(output)-1][0], val[1])
		} else {
			output = append(output, val)
		}
	}
	return output
}
