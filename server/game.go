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
func Play(board [][]int, player int, row int, col int) bool {
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

	// do two passes of the connected component
	//   1. determine if no liberties remaing
	//   2. remove stones if none remaining

	// note that we need to do step 1 for all directions first
	// since captured stones may still capture other stones
	up := CanCapture(board, row-1, col)
	down := CanCapture(board, row+1, col)
	left := CanCapture(board, row, col-1)
	right := CanCapture(board, row, col+1)

	if up {
		Capture(board, row-1, col)
	}
	if down {
		Capture(board, row-1, col)
	}
	if left {
		Capture(board, row, col-1)
	}
	if right {
		Capture(board, row, col+1)
	}

	// Optional Rule 7A: do not allow self captures
	// because that would be silly and also the game would never end? Would be funny tho
	return true
}
func CanCapture(board [][]int, row int, col int) bool {
	if row < 0 || row > len(board)-1 || col < 0 || col > len(board[row])-1 {
		return false
	}
	player := board[row][col]

	var DFS func(row int, col int) bool
	DFS = func(row int, col int) bool {
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
		captured := DFS(row-1, col) && DFS(row+1, col) && DFS(row, col-1) && DFS(row, col+1)
		board[row][col] = player
		return captured
	}
	return DFS(row, col)
}

func Capture(board [][]int, row int, col int) {
	player := board[row][col]

	var DFS func(row int, col int)
	DFS = func(row int, col int) {
		if row < 0 || row > len(board)-1 || col < 0 || col > len(board[row])-1 {
			return
		}
		if board[row][col] != player {
			return
		}
		board[row][col] = 0
		DFS(row-1, col)
		DFS(row+1, col)
		DFS(row, col-1)
		DFS(row, col+1)
	}
	DFS(row, col)
}
