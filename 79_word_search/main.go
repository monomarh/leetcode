package main

import "fmt"

func main() {
	grid := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}

	fmt.Println(exist(grid, "ABCCED"))
}

func exist(board [][]byte, word string) bool {
	lenX, lenY := len(board), len(board[0])

	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			if board[i][j] == word[0] && dig(board, i, j, 0, word) {
				return true
			}
		}
	}

	return false
}

func dig(board [][]byte, i, j, count int, word string) bool {
	if len(word) == count {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[count] {
		return false
	}

	temp := board[i][j]
	board[i][j] = '*'

	found := dig(board, i+1, j, count+1, word) ||
		dig(board, i, j+1, count+1, word) ||
		dig(board, i-1, j, count+1, word) ||
		dig(board, i, j-1, count+1, word)

	board[i][j] = temp

	return found
}
