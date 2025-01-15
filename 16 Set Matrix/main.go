package main

func main() {

}

func setZeroesBruteforce() {
	// take array for row and column mark if there is any zero
	// iterate the array and modify the array accordingly

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
