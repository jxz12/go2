package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard(5)

	require.Equal(t, 5, board.Width(), "board does not have 5 rows")

	placed := board.Play(1, 0, 0)
	assert.True(t, placed == true && board.Get(0, 0) == 1, "could not place stone")
	placed = board.Play(1, 0, 0)
	assert.False(t, placed, "placed stone in occupied position")
}

func AssertBoardsEqual(t *testing.T, observed Board, expected Board) {
	assert.Equal(t, expected, observed, "capture did not remove pieces correctly")
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
	assert.Falsef(t, board.Play(1, 1, 1), "should not be able to play move due to self-capture\n%s", board)
}

func TestPrint(t *testing.T) {
	board := NewBoard(5)
	assert.Equal(t, board.String(), "0 0 0 0 0\n0 0 0 0 0\n0 0 0 0 0\n0 0 0 0 0\n0 0 0 0 0", "empty board not printed correctly")
}

func TestScore(t *testing.T) {
	board := Board{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	scores := board.Score()
	assert.Equalf(t, len(scores), 0, "score is not empty: %v", scores)
	board = Board{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	scores = board.Score()
	assert.Truef(t, len(scores) == 1 && scores[1] == 25, "board is not owned by 1: %v", scores)
	board = Board{
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{1, 1, 1, 0, 0},
		{0, 0, 0, 2, 0},
		{0, 0, 0, 0, 0},
	}
	scores = board.Score()
	assert.Truef(t, len(scores) == 2 && scores[1] == 9 && scores[2] == 1, "score is not correct: %v", scores)
}
