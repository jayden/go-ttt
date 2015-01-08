package tictactoe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDrawGame(t *testing.T) {
	board := NewBoard()
	t.Log("Empty board should not be a draw")
	assert.False(t, GameDraw(board))
	t.Log("Full board should result in a draw")
	assert.True(t, GameDraw(tieBoard()))
	t.Log("A full board with a winner should not result in a draw")
	assert.False(t, GameDraw(fullWinBoard()))
}

func tieBoard() *Board {
	tieBoard := NewBoard()
	FillSpaces(tieBoard, "X", 0, 1, 5, 6, 7)
	FillSpaces(tieBoard, "O", 2, 3, 4, 8)
	return tieBoard
}

func fullWinBoard() *Board {
	fullWinBoard := tieBoard()
	fullWinBoard.FillSpace(2, "X")
	return fullWinBoard
}

func winBoard() *Board {
	winBoard := NewBoard()
	FillSpaces(winBoard, "X", 0, 1, 2)
	return winBoard
}

func altWinBoard() *Board {
	altWinBoard := NewBoard()
	FillSpaces(altWinBoard, "O", 3, 4, 5)
	return altWinBoard
}

func colWinBoard() *Board {
	colWinBoard := NewBoard()
	FillSpaces(colWinBoard, "X", 0, 3, 6)
	return colWinBoard
}

func inProgressBoard() *Board {
	inProgressBoard := NewBoard()
	inProgressBoard.FillSpace(0, "X")
	inProgressBoard.FillSpace(4, "O")
	return inProgressBoard
}

func TestWinGame(t *testing.T) {
	board := NewBoard()
	t.Log("Empty board should not be a win")
	_, gameWon := GameWon(board)
	assert.False(t, gameWon)

	t.Log("Tie board should not be a win")
	_, gameWon = GameWon(tieBoard())
	assert.False(t, gameWon)

	t.Log("Three in a row results in a win")
	winner, gameWon := GameWon(winBoard())
	assert.True(t, gameWon)

	t.Log("Three in another row results in a win")
	_, gameWon = GameWon(altWinBoard())
	assert.True(t, gameWon)

	t.Log("A column match is also a win")
	_, gameWon = GameWon(colWinBoard())
	assert.True(t, gameWon)

	t.Log("Gets winner of game")
	assert.Equal(t, "X", winner)
}

func TestGetAvailableMoves(t *testing.T) {
	emptyBoard := NewBoard()
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	t.Log("Empty board should return all spaces")
	assert.Equal(t, expected, GetAvailableMoves(emptyBoard))

	t.Log("Full board should return no spaces")
	assert.Equal(t, []int{}, GetAvailableMoves(tieBoard()))

	t.Log("A winning board should also return no spaces")
	assert.Equal(t, []int{}, GetAvailableMoves(winBoard()))

	t.Log("An in-progress board should return available spaces")
	expected = []int{1, 2, 3, 5, 6, 7, 8}
	assert.Equal(t, expected, GetAvailableMoves(inProgressBoard()))
}
