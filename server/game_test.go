package main

import "testing"
import "reflect"

func TestNewBoard(t *testing.T) {
	board := NewBoard(5)

	if len(board) != 5 {
		t.Errorf("board does not have 5 rows")
	}
	if len(board[0]) != 5 {
		t.Errorf("board does not have 5 columns")
	}

	placed := Play(board, 1, 0, 0)
	if placed != true || board[0][0] != 1 {
		t.Errorf("could not place stone")
	}
	placed = Play(board, 1, 0, 0)
	if placed == true {
		t.Errorf("placed stone in occupied position")
	}
	scores := Score(board)
	if scores[1] != 1 || scores[0] != 24 {
		t.Errorf("score is not 1=1 and 0=24")
	}
}

func TestPlay(t *testing.T) {
	board := [][]int{
		{1, 1, 2, 0, 0},
		{1, 3, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{2, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	Play(board, 3, 2, 1)
	expected := [][]int{
		{0, 0, 2, 0, 0},
		{0, 3, 0, 0, 0},
		{0, 3, 0, 0, 0},
		{2, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	if !reflect.DeepEqual(board, expected) {
		t.Errorf("capture did not remove pieces correctly, got:\n%s\nexpected:\n%s", ToString(board), ToString(expected))
	}
}

func TestPrint(t *testing.T) {
	board := NewBoard(5)
	if ToString(board) != "0 0 0 0 0\n0 0 0 0 0\n0 0 0 0 0\n0 0 0 0 0\n0 0 0 0 0" {
		t.Errorf("empty board not printed correctly")
	}
}
