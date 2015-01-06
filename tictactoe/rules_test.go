package tictactoe

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestDrawGame(t *testing.T) {
  board := MakeBoard()
  t.Log("Empty board should not be a draw")
  assert.False(t, board.GameDraw())
  t.Log("Full board should result in a draw")
  assert.True(t, tieBoard().GameDraw())
  t.Log("A full board with a winner should not result in a draw")
  assert.False(t, fullWinBoard().GameDraw())
}

func tieBoard() *Board {
  tieBoard := MakeBoard()
  FillSpaces(tieBoard, "X", 0, 1, 5, 6, 7)
  FillSpaces(tieBoard, "O", 2, 3, 4, 8)
  return tieBoard
}

func fullWinBoard() *Board {
  fullWinBoard := tieBoard()
  fullWinBoard.FillSpace(2, "X")
  return fullWinBoard
}

func winBoard() *Board {
  winBoard := MakeBoard()
  FillSpaces(winBoard, "X", 0, 1, 2)
  return winBoard
}

func altWinBoard() *Board {
  altWinBoard := MakeBoard()
  FillSpaces(altWinBoard, "O", 3, 4, 5)
  return altWinBoard
}

func colWinBoard() *Board {
  colWinBoard := MakeBoard()
  FillSpaces(colWinBoard, "X", 0, 3, 6)
  return colWinBoard
}

func inProgressBoard() *Board {
  inProgressBoard := MakeBoard()
  inProgressBoard.FillSpace(0, "X")
  inProgressBoard.FillSpace(4, "O")
  return inProgressBoard
}

func TestWinGame(t *testing.T) {
  board := MakeBoard()
  t.Log("Empty board should not be a win")
  _, gameWon := board.GameWon()
  assert.False(t, gameWon)

  t.Log("Tie board should not be a win")
  _, gameWon = tieBoard().GameWon()
  assert.False(t, gameWon)

  t.Log("Three in a row results in a win")
  winner, gameWon := winBoard().GameWon()
  assert.True(t, gameWon)

  t.Log("Three in a row results in a win (2)")
  _, gameWon = altWinBoard().GameWon()
  assert.True(t, gameWon)

  t.Log("A column match is also a win")
  _, gameWon = colWinBoard().GameWon()
  assert.True(t, gameWon)

  t.Log("Gets winner of game")
  assert.Equal(t, "X", winner)
}

func TestGetAvailableMoves(t *testing.T) {
  emptyBoard := MakeBoard()
  expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
  t.Log("Empty board should return all spaces")
  assert.Equal(t, expected, emptyBoard.GetAvailableMoves())

  t.Log("Full board should return no spaces")
  assert.Equal(t, []int{}, tieBoard().GetAvailableMoves())

  t.Log("A winning board should also return no spaces")
  assert.Equal(t, []int{}, winBoard().GetAvailableMoves())

  t.Log("An in-progress board should return available spaces")
  expected = []int{1, 2, 3, 5, 6, 7, 8}
  assert.Equal(t, expected, inProgressBoard().GetAvailableMoves())
}
