package tictactoe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupNewGame() *Game {
	board := NewBoard()
	players := []Player{new(MockPlayer), new(MockPlayer)}
	ui := new(MockConsoleUI)
	return NewGame(board, players, ui)
}

func TestGameInitializesNewBoard(t *testing.T) {
	game := setupNewGame()
	emptyBoard := NewBoard()

	t.Log("Game has an empty board")
	assert.Equal(t, emptyBoard, game.board)
}

func TestGamePutsMarkerOnBoard(t *testing.T) {
	board := NewBoard()
	board.FillSpace(0, x)
	game := setupNewGame()

	game.PutMove(0, x)

	t.Log("Game puts marker on the board")
	assert.Equal(t, board, game.board)
}

func TestGameRunner(t *testing.T) {
	board := NewBoard()
	player1, player2 := new(MockPlayer), new(MockPlayer)
	players := []Player{player1, player2}
	console := new(MockConsoleUI)
	game := NewGame(board, players, console)

	console.MakeFakeGameSelection(humanFirst)

	t.Log("Player 1 gets last move if human is first")
	FillSpaces(board, player1.Marker(), 0, 2, 5, 7)
	FillSpaces(board, player2.Marker(), 1, 3, 4, 8)
	player1.MakeFakeMove(6)
	game.Run()
	assert.Equal(t, board.Spaces()[6], player1.Marker())

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
	assert.Equal(t, board.Spaces()[6], player2.Marker())

	console.MakeFakeGameSelection(exit)
	console.ResetCall()
	ClearBoard(board)

	t.Log("Game exits")
	game.Run()
	assert.False(t, console.Called("PrintBoard"))

}
