package tictactoe

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestMakeBlankBoard(t *testing.T) {
  actual := MakeBoard().spaces
  expected := []string{blank, blank, blank, blank, blank, blank, blank, blank, blank}

  t.Log("can create a blank board")
  assert.Equal(t, expected, actual)
}

func TestFillBoard(t *testing.T) {
  board := MakeBoard()
  board.FillSpace(4, "X")

  t.Log("can fill a board with a marker")
  assert.Equal(t, board.spaces[4], "X")
}
/*
func TestFillIgnoresInvalidPositions(t *testing.T) {
  board := MakeBoard()
  copyBoard := make([]string, len(board.spaces))
  copy(copyBoard, board.spaces)
  board.FillSpace(9, "X")

  t.Log("board should not be filled with an invalid position")
  assert.Equal(t, board.spaces, copyBoard)
}*/

func TestClearSpace(t *testing.T) {
  board := MakeBoard()
  board.FillSpace(4, "X")

  t.Log("can clear a space on the board")
  assert.Equal(t, board.spaces[4], "X")
  board.ClearSpace(4)
  assert.Equal(t, board.spaces[4], " ")
}
