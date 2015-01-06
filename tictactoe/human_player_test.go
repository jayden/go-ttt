package tictactoe

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestHumanIsAPlayer(t *testing.T) {
  human := new(Human)
  var player Player = human

  t.Log("Human implements the Player interface")
  assert.Equal(t, human, player)
}

func TestManipulateHumanMark(t *testing.T) {
  human := new(Human)
  t.Log("gets and sets the marker of human player")
  human.SetMarker("X")
  assert.Equal(t, "X", human.GetMarker())

  t.Log("changes the marker of human player")
  human.SetMarker("O")
  assert.Equal(t, "O", human.GetMarker())

  t.Log("human player's opponent marker should be its opposite")
  assert.Equal(t, "X", human.GetOpponentMarker())
}

func TestMakesMove(t *testing.T) {
  human := new(Human)
  mockMove := "1"
  writeMockInput(mockMove)
  board := MakeBoard()

  t.Log("captures user move selection")
  human.GetMove(board)
  assert.Equal(t, 1, human.move)
}