package main

import "testing"
import "reflect"

func TestNewBoard(t *testing.T) {
	board := NewBoard(5)

	if board.Width() != 5 {
		t.Errorf("board does not have 5 rows")
	}

	placed := board.Play(1, 0, 0)
	if placed != true || board.Get(0, 0) != 1 {
		t.Errorf("could not place stone")
	}
	placed = board.Play(1, 0, 0)
	if placed == true {
		t.Errorf("placed stone in occupied position")
	}
	scores := board.Score()
	if scores[1] != 1 || scores[0] != 24 {
		t.Errorf("score is not 1=1 and 0=24")
	}
}

func TestPlay(t *testing.T) {
	board := Board{
		{1, 1, 2, 0, 0},
		{1, 3, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{2, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	board.Play(3, 2, 1)
	expected := Board{
		{0, 0, 2, 0, 0},
		{0, 3, 0, 0, 0},
		{0, 3, 0, 0, 0},
		{2, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	AssertBoardsEqual(t, board, expected)

	board = Board{
		{1, 1, 1, 3, 0},
		{1, 1, 1, 2, 3},
		{1, 1, 0, 2, 3},
		{1, 2, 2, 2, 3},
		{2, 3, 3, 3, 0},
	}
	board.Play(3, 2, 2)
	expected = Board{
		{0, 0, 0, 3, 0},
		{0, 0, 0, 0, 3},
		{0, 0, 3, 0, 3},
		{0, 0, 0, 0, 3},
		{2, 3, 3, 3, 0},
	}
	AssertBoardsEqual(t, board, expected)
}

func AssertBoardsEqual(t *testing.T, observed Board, expected Board) {
	if !reflect.DeepEqual(observed, expected) {
		t.Errorf("capture did not remove pieces correctly, observed:\n%s\nexpected:\n%s", observed.ToString(), expected.ToString())
	}
}

func TestPrint(t *testing.T) {
	board := NewBoard(5)
	if board.ToString() != "0 0 0 0 0\n0 0 0 0 0\n0 0 0 0 0\n0 0 0 0 0\n0 0 0 0 0" {
		t.Errorf("empty board not printed correctly")
	}
}
