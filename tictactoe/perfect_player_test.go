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
  assert.Equal(t, "O", perfectPlayer.GetMarker())

  t.Log("PerfectPlayer's opponent marker is its opposite")
  assert.Equal(t, "X", perfectPlayer.GetOpponentMarker())
}

func TestPerfectPlayerGetsBestMoves(t *testing.T) {
  perfectPlayer := new(PerfectPlayer)
  perfectPlayer.SetMarker("O")

  //t.Log("PerfectPlayer should return winning move")
  //assert.Equal(t, 2, perfectPlayer.GetMove(testBoard()))

  t.Log("PerfectPlayer should block winning move")
  assert.Equal(t, 8, perfectPlayer.GetMove(blockBoard()))
}

func blockBoard() *Board {
  board := MakeBoard()
  board.Fill(0, "X")
  board.Fill(1, "O")
  board.Fill(3, "O")
  board.Fill(4, "O")
  board.Fill(5, "X")
  board.Fill(6, "X")
  board.Fill(7, "X")
  return board
}

func testBoard() *Board {
  board := MakeBoard()
  board.Fill(0, "X")
  board.Fill(4, "O")
  board.Fill(7, "X")
  board.Fill(6, "O")
  board.Fill(3, "X")
  return board
}
