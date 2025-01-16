package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))

	matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix))

}

func spiralOrder(matrix [][]int) []int {
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	result := make([]int, 0)

	for top <= bottom && left <= right {

		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}

	}
	return result
}
