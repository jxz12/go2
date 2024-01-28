package main

import "bytes"
import "strconv"
import "fmt"

func NewBoard(size int) [][]int {
	board := make([][]int, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
	}
	return board
}

func ToString(board [][]int) string {
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

func Score(board [][]int) map[int]int {
	scores := make(map[int]int)
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			scores[board[row][col]] += 1
		}
	}
	return scores
}
func Place(board [][]int, player int, row int, col int) bool {
	if player == 0 {
		return false // 0 cannot be a player id
	}
	// Rule 7 in https://en.wikipedia.org/wiki/Rules_of_Go
	// Step 1: position must be empty
	if board[row][col] != 0 {
		return false
	}
	board[row][col] = player
	// Step 2: look for groups of stones with liberties gone
	// for each adjacent stone, try doing a BFS on that colour to check if all connecting liberties are gone
	// remove stones if all gone

	// Optional Rule 7A: do not allow self captures
	// because that would be silly and also the game would never end? Would be funny tho
	return true
}

func Capture(board [][]int, row int, col int) bool {
	player := board[row][col]

	// do two passes of the connected component
	//   1. determine if no liberties remaing
	//   2. remove stone if none remaining
	var CanCapture func(row int, col int) bool
	CanCapture = func(row int, col int) bool {
		if row < 0 || row > len(board)-1 || col < 0 || col > len(board[row])-1 {
			return true
		}
		if board[row][col] == 0 {
			// empty square means liberty is available
			return false
		}
		if board[row][col] != player {
			// liberty taken by another colour stone
			return true
		}
		// mark position as explored by setting to minus 1
		board[row][col] = -player
		captured = CanCapture(row-1, col) && CanCapture(row+1, col) && CanCapture(row, col-1) && CanCapture(row, col+1)
		board[row][col] = player
		return captured
	}
	captured := CanCapture(row, col)
	return captured
}
