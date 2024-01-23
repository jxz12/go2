package main

func NewBoard(size int) [][]int {
	board := make([][]int, size)
	for i:=0; i<size; i++ {
		board[i] = make([]int, size)
	}
	return board
}

func Place(board [][]int, row int, col int, player int) bool {
	if board[row][col] != 0 {
		return false
	}
	board[row][col] = player
	return true
}

func Score(board [][]int) map[int]int{
	scores := make(map[int]int)
	for row:=0; row<len(board); row++ {
		for col:=0; col<len(board[row]); col++ {
			scores[board[row][col]] += 1
		}
	}
	return scores
}
