package main

import "bytes"
import "strconv"

func NewBoard(size int) [][]int {
	board := make([][]int, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
	}
	return board
}

func PrintBoard(board [][]int) string {
	buffer := bytes.NewBufferString("")

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			buffer.WriteString(strconv.Itoa(board[row][col]))
			if col < len(board[row])-1 {
				buffer.WriteString(" ")
			}
		}
		if row < len(board)-1 {
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}

func Place(board [][]int, player int, row int, col int) bool {
	if board[row][col] != 0 {
		return false
	}
	board[row][col] = player
	return true
}

func Score(board [][]int) map[int]int {
	scores := make(map[int]int)
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			scores[board[row][col]] += 1
		}
	}
	return scores
}
