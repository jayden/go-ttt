package tictactoe

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestPerfectPlayerIsAPlayer(t *testing.T) {
  perfectPlayer := new(PerfectPlayer)
  var player Player = perfectPlayer

  t.Log("PerfectPlayer implements the Player interface")
  assert.True(t, player == perfectPlayer)
}

func TestPerfectPlayerCanManipulateMarker(t *testing.T) {
  perfectPlayer := new(PerfectPlayer)
  perfectPlayer.SetMarker("O")

  t.Log("PerfectPlayer can set its mark")
  assert.Equal(t, "O", perfectPlayer.Marker())
}

func TestPerfectPlayerGetsBestMoves(t *testing.T) {
  perfectPlayer := new(PerfectPlayer)
  perfectPlayer.SetMarker("O")

  t.Log("PerfectPlayer should return winning move")
  assert.Equal(t, 2, perfectPlayer.Move(testBoard()))

  t.Log("PerfectPlayer should block winning move")
  assert.Equal(t, 8, perfectPlayer.Move(blockBoard()))

  t.Log("PerfectPlayer should play center space")
  firstMoveBoard := MakeBoard()
  firstMoveBoard.FillSpace(0, "X")
  assert.Equal(t, 4, perfectPlayer.Move(firstMoveBoard))

  //t.Log("PerfectPlayer plays corner spaces if first")
  //blankBoard := MakeBoard()
  //perfectPlayer.SetMarker("X")
  //corners := []int{0, 2, 6, 8}
  //move := perfectPlayer.GetMove(blankBoard)
  //assert.True(t, valueInSlice(move, corners))
}

func valueInSlice(value int, slice []int) bool {
  for _,v := range slice {
    if v == value { return true }
  }
  return false
}

func blockBoard() *Board {
  board := MakeBoard()
  FillSpaces(board, "X", 0, 5, 6, 7)
  FillSpaces(board, "O", 1, 3, 4)
  return board
}

func testBoard() *Board {
  board := MakeBoard()
  FillSpaces(board, "X", 0, 3, 7)
  FillSpaces(board, "O", 4, 6)
  return board
}
