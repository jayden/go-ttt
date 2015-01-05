package tictactoe

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestManipulateHumanMark(t *testing.T) {
  human := new(Human)
  t.Log("gets and sets the marker of human player")
  human.SetMarker("X")
  assert.Equal(t, "X", human.GetMarker())

  t.Log("changes the marker of human player")
  human.SetMarker("O")
  assert.Equal(t, "O", human.GetMarker())
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
