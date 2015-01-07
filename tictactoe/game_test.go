package tictactoe

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func setupNewGame() *Game {
  board := MakeBoard()
  players := []Player{ new(MockPlayer), new(MockPlayer) }
  ui := new(MockConsoleUI)
  return NewGame(board, players, ui)
}

func TestGameInitializesNewBoard(t *testing.T) {
  game := setupNewGame()
  emptyBoard := MakeBoard()

  t.Log("Game has an empty board")
  assert.Equal(t, emptyBoard, game.Board())
}

func TestGamePutsMarkerOnBoard(t *testing.T) {
  board := MakeBoard()
  board.FillSpace(0, "X")
  game := setupNewGame()

  game.PutMove(0, "X")

  t.Log("Game puts marker on the board")
  assert.Equal(t, board, game.Board())
}

func TestGameValidatesMoves(t *testing.T) {
  game := setupNewGame()

  t.Log("Game rejects invalid moves")
  assert.Equal(t, false, game.IsValidMove("-4"))
  assert.Equal(t, false, game.IsValidMove("m"))
}
