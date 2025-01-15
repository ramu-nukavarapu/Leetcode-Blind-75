package main

import "fmt"

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)

	matrix = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroesBruteforce() {
	// take array's for row and column marking if there is any zero in the matrix
	// iterate the array based on the row and column marking modify the array accordingly
}

func setZeroes(matrix [][]int) {
	col := 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				if j != 0 {
					matrix[0][j] = 0
				} else {
					col = 0
				}
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	if col == 0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
