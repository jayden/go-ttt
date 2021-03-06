package tictactoe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerfectPlayerIsAPlayer(t *testing.T) {
	perfectPlayer := new(PerfectPlayer)
	var player Player = perfectPlayer

	t.Log("PerfectPlayer implements the Player interface")
	assert.True(t, player == perfectPlayer)
}

func TestPerfectPlayerCanManipulateMarker(t *testing.T) {
	perfectPlayer := new(PerfectPlayer)
	perfectPlayer.SetMarker(o)

	t.Log("PerfectPlayer can set its mark")
	assert.Equal(t, o, perfectPlayer.Marker())
}

func TestPerfectPlayerGetsBestMoves(t *testing.T) {
	perfectPlayer := new(PerfectPlayer)
	perfectPlayer.SetMarker(o)

	t.Log("PerfectPlayer should return winning move")
	assert.Equal(t, 2, perfectPlayer.Move(testBoard()))

	t.Log("PerfectPlayer should block winning move")
	assert.Equal(t, 8, perfectPlayer.Move(blockBoard()))

	t.Log("PerfectPlayer should play center space")
	firstMoveBoard := NewBoard()
	firstMoveBoard.FillSpace(0, x)
	assert.Equal(t, 4, perfectPlayer.Move(firstMoveBoard))
}

func blockBoard() *Board {
	board := NewBoard()
	FillSpaces(board, x, 0, 5, 6, 7)
	FillSpaces(board, o, 1, 3, 4)
	return board
}

func testBoard() *Board {
	board := NewBoard()
	FillSpaces(board, x, 0, 3, 7)
	FillSpaces(board, o, 4, 6)
	return board
}
