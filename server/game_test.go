package main

import (
	"fmt"
	"reflect"
	"testing"
)

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

func TestCapture(t *testing.T) {
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

	board = Board{
		{0, 1, 2, 0, 0},
		{1, 2, 0, 0, 0},
		{2, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	board.Play(2, 0, 0)
	expected = Board{
		{2, 0, 2, 0, 0},
		{0, 2, 0, 0, 0},
		{2, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	AssertBoardsEqual(t, board, expected)
}

func TestNoSelfCapture(t *testing.T) {
	board := Board{
		{1, 1, 1, 2, 0},
		{1, 0, 1, 2, 0},
		{1, 1, 2, 0, 0},
		{2, 2, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	if board.Play(1, 1, 1) {
		t.Errorf("should not be able to play move due to self-capture\n%s", board.ToString())
	}
}

func TestWalrus(t *testing.T) {
	// err := "error"
	// fmt.Println(err)
	// if err := "error2"; err == "error2" {
	// 	fmt.Println(err)
	// }
	var x int
	x = 5
	px := &x
	fmt.Println(px)
	ppx := &px
	fmt.Println(ppx)
	fmt.Println(**ppx)
}

func TestBigArray(t *testing.T) {
	var bigArray [10000]int
	WriteBigArray(bigArray)
	var i int
	i = 1
	i = -(-(-(-(-(-(-(-(-(-(-(-(-(-(-i))))))))))))))
	t.Errorf("%v", bigArray[0])
}
func WriteBigArray(bigArray [10000]int) {
	bigArray[0] = 1
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
