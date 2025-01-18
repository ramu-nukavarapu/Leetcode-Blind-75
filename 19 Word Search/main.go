package main

import "fmt"

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))

	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word = "SEE"
	fmt.Println(exist(board, word))

	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word = "ABCB"
	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}

	var isWord func(x, y int, word string) bool
	isWord = func(x, y int, word string) bool {
		if len(word) == 0 {
			return true
		}
		if x < 0 || x >= len(board) ||
			y < 0 || y >= len(board[0]) ||
			board[x][y] != word[0] {
			return false
		}
		tmp := board[x][y]
		board[x][y] = '#'
		res := isWord(x+1, y, word[1:]) ||
			isWord(x-1, y, word[1:]) ||
			isWord(x, y+1, word[1:]) ||
			isWord(x, y-1, word[1:])
		board[x][y] = tmp
		return res
	}

	for i, row := range board {
		for j := range row {
			if isWord(i, j, word) {
				return true
			}
		}
	}

	return false
}
