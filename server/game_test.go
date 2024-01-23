package main

import "testing"
import "fmt"

func TestPlacement(t *testing.T) {
	board := NewBoard(5)

	if len(board) != 5 {
		t.Errorf("board does not have 5 rows")
	}
	if len(board[0]) != 5 {
		t.Errorf("board does not have 5 columns")
	}

	placed := Place(board, 0, 0, 1)
	if placed != true || board[0][0] != 1 {
		t.Errorf("could not place stone")
	}
	placed = Place(board, 0, 0, 1)
	if placed == true {
		t.Errorf("placed stone in occupied position")
	}
	scores := Score(board)
	if scores[1] != 1 || scores[0] != 24 {
		t.Errorf("score is not 1=1 and 0=24")
	}
}

func TestPrint(t *testing.T) {
	board := NewBoard(5)
	fmt.Println(PrintBoard(board))
}
