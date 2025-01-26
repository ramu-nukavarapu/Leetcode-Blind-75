package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{1, 3}, {6, 9}, {2, 5}}
	fmt.Println(eraseOverlapIntervals(arr))

	arr = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}, {4, 8}}
	fmt.Println(eraseOverlapIntervals(arr))

	arr = [][]int{{2, 3}, {6, 9}, {0, 1}}
	fmt.Println(eraseOverlapIntervals(arr))

	arr = [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println(eraseOverlapIntervals(arr))
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := 0
	prevEnd := intervals[0][1]

	for _, val := range intervals[1:] {
		if val[0] >= prevEnd {
			prevEnd = val[1]
		} else {
			result++
			prevEnd = min(prevEnd, val[1])
		}
	}
	return result
}
