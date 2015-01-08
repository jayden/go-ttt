package tictactoe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockPlayerIsAPlayer(t *testing.T) {
	mockPlayer := new(MockPlayer)
	var player Player = mockPlayer

	t.Log("Mock Player is a kind of Player")
	assert.Equal(t, mockPlayer, player)
}

func TestMockPlayerManipulatesMarker(t *testing.T) {
	mockPlayer := new(MockPlayer)
	mockPlayer.SetMarker("X")

	t.Log("Mock Player can get and set marker")
	assert.Equal(t, "X", mockPlayer.Marker())
}

func TestMockPlayerMakesFakeMoves(t *testing.T) {
	mockPlayer := new(MockPlayer)
	mockPlayer.MakeFakeMove(0)
	board := NewBoard()

	t.Log("Mock Player makes fake moves")
	assert.Equal(t, 0, mockPlayer.Move(board))
}
