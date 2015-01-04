package tictactoe

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestMakeBlankBoard(t *testing.T) {
  actual := MakeBoard().spaces
  expected := []string{blank, blank, blank, blank, blank, blank, blank, blank, blank}
  assert.Equal(t, expected, actual)
}

func TestFillBoard(t *testing.T) {
  board := MakeBoard()
  board.Fill(4, "X")
  assert.Equal(t, board.spaces[4], "X")
}

func TestFillIgnoresInvalidPositions(t *testing.T) {
  board := MakeBoard()
  fakeBoard := make([]string, len(board.spaces))
  copy(fakeBoard, board.spaces)
  board.Fill(9, "X")
  assert.Equal(t, board.spaces, fakeBoard)
}

func TestClearSpace(t *testing.T) {
  board := MakeBoard()
  board.Fill(4, "X")
  assert.Equal(t, board.spaces[4], "X")
  board.ClearSpace(4)
  assert.Equal(t, board.spaces[4], " ")
}
