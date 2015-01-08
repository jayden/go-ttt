package tictactoe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHumanIsAPlayer(t *testing.T) {
	human := NewHumanPlayer()
	var player Player = human

	t.Log("Human implements the Player interface")
	assert.Equal(t, human, player)
}

func TestManipulateHumanMark(t *testing.T) {
	human := NewHumanPlayer()
	t.Log("gets and sets the marker of human player")
	human.SetMarker("X")
	assert.Equal(t, "X", human.Marker())

	t.Log("changes the marker of human player")
	human.SetMarker("O")
	assert.Equal(t, "O", human.Marker())
}

func TestMakesMove(t *testing.T) {
	human := NewHumanPlayer()
	mockMove := "1"
	writeMockInput(mockMove)
	board := NewBoard()

	t.Log("captures user move selection")
	move := human.Move(board)
	assert.Equal(t, 1, move)
}
