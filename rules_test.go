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

func TestWinGame(t *testing.T) {
  board := MakeBoard()
  t.Log("Empty board should not be a win")
  assert.False(t, board.GameWon())

  t.Log("Tie board should not be a win")
  assert.False(t, tieBoard().GameWon())

  t.Log("Three in a row results in a win")
  assert.True(t, winBoard().GameWon())

  t.Log("Three in a row results in a win (2)")
  assert.True(t, altWinBoard().GameWon())

  t.Log("A column match is also a win")
  assert.True(t, colWinBoard().GameWon())
}

func TestGetAvailableMoves(t *testing.T) {
  emptyBoard := MakeBoard()
  expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
  t.Log("Empty board should return all spaces")
  assert.Equal(t, expected, emptyBoard.GetAvailableMoves())

  t.Log("Full board should return no spaces")
  assert.Equal(t, []int{}, tieBoard().GetAvailableMoves())
}
