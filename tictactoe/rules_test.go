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
  tieBoard.Fill(0, "X")
  tieBoard.Fill(1, "X")
  tieBoard.Fill(2, "O")
  tieBoard.Fill(3, "O")
  tieBoard.Fill(4, "O")
  tieBoard.Fill(5, "X")
  tieBoard.Fill(6, "X")
  tieBoard.Fill(7, "X")
  tieBoard.Fill(8, "O")
  return tieBoard
}

func fullWinBoard() *Board {
  fullWinBoard := tieBoard()
  fullWinBoard.Fill(2, "X")
  return fullWinBoard
}

func winBoard() *Board {
  winBoard := MakeBoard()
  winBoard.Fill(0, "X")
  winBoard.Fill(1, "X")
  winBoard.Fill(2, "X")
  return winBoard
}

func altWinBoard() *Board {
  altWinBoard := MakeBoard()
  altWinBoard.Fill(3, "O")
  altWinBoard.Fill(4, "O")
  altWinBoard.Fill(5, "O")
  return altWinBoard
}

func colWinBoard() *Board {
  colWinBoard := MakeBoard()
  colWinBoard.Fill(0, "X")
  colWinBoard.Fill(3, "X")
  colWinBoard.Fill(6, "X")
  return colWinBoard
}

func inProgressBoard() *Board {
  inProgressBoard := MakeBoard()
  inProgressBoard.Fill(0, "X")
  inProgressBoard.Fill(4, "O")
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
