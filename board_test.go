package tictactoe

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

const blank = " "

func TestMakeEmptyBoard(t *testing.T) {
  board := MakeBoard()
  spaces := board.spaces
  expected := []string{blank, blank, blank, blank, blank, blank, blank, blank, blank}
  assert.Equal(t, expected, spaces)
}
