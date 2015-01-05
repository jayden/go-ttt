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
  human.SetMarker("X")
  board := MakeBoard()
  t.Log("makes a move on the board")
  human.GetMove(board)
  expected := []string{"X", blank, blank, blank, blank, blank, blank, blank, blank}
  assert.Equal(t, expected, board.spaces)
}
