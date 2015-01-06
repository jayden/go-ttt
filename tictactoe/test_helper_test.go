package tictactoe

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestFillMultipleSpacesAtOnce(t *testing.T) {
  t.Log("helper method #FillSpaces fills multiple spaces at once")
  board := MakeBoard()
  FillSpaces(board, "X", 0, 1, 2)
  expected := []string{"X", "X", "X", blank, blank, blank, blank, blank, blank}
  assert.Equal(t, expected, board.spaces)
}
