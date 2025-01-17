package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotateImageOptimal(matrix)
	fmt.Println(matrix)

	matrix = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotateImageOptimal(matrix)
	fmt.Println(matrix)
}

func rotateImageBruteforce(matrix [][]int) {
	// take another array wit n X n size
	// Iterate the input array from last add the elements column wise into the array
}

func rotateImageOptimal(matrix [][]int) {
	var row, col, temp int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if i >= j {
				continue
			}
			temp = matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}

	for i := 0; i < len(matrix); i++ {
		row = 0
		col = len(matrix) - 1
		for row < col {
			temp = matrix[i][row]
			matrix[i][row] = matrix[i][col]
			matrix[i][col] = temp
			row++
			col--
		}
	}

}
