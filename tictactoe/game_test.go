package tictactoe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupNewGame() *Game {
	board := MakeBoard()
	players := []Player{new(MockPlayer), new(MockPlayer)}
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

func TestGameRunner(t *testing.T) {
	board := MakeBoard()
	player1, player2 := new(MockPlayer), new(MockPlayer)
	players := []Player{player1, player2}
	console := new(MockConsoleUI)
	player1.SetMarker("X")
	player2.SetMarker("O")
	game := Game{board, players, console}

	console.MakeFakeGameSelection(humanFirst)

	t.Log("Player 1 gets last move if human is first")
	FillSpaces(board, player1.Marker(), 0, 2, 5, 7)
	FillSpaces(board, player2.Marker(), 1, 3, 4, 8)
	player1.MakeFakeMove(6)
	game.Run()
	assert.Equal(t, board.spaces[6], player1.Marker())

	t.Log("Prints game conclusion at end of game")
	assert.True(t, console.Called("PrintGameConclusion"))

	console.ResetCall()
	ClearBoard(board)

	t.Log("Prints game conclusion right when there is a winner")
	FillSpaces(board, player1.Marker(), 0, 1)
	FillSpaces(board, player2.Marker(), 3, 4)
	player1.MakeFakeMove(2)
	game.Run()
	assert.True(t, console.Called("PrintGameConclusion"))

	console.MakeFakeGameSelection(computerFirst)
	console.ResetCall()
	ClearBoard(board)

	t.Log("Player 2 gets last move if computer is first")
	FillSpaces(board, player1.Marker(), 0, 2, 5, 7)
	FillSpaces(board, player2.Marker(), 1, 3, 4, 8)
	player2.MakeFakeMove(6)
	game.Run()
	assert.Equal(t, board.spaces[6], player2.Marker())

	console.MakeFakeGameSelection(exit)
	console.ResetCall()
	ClearBoard(board)

	t.Log("Game exits")
	game.Run()
	assert.False(t, console.Called("PrintBoard"))

}
