package main

import "github.com/jayden/go-tictactoe/tictactoe"

func main() {
	board := tictactoe.NewBoard()
	ui := tictactoe.NewConsoleUI()
	player1 := tictactoe.NewHumanPlayer()
	player2 := tictactoe.NewPerfectPlayer()
	players := []tictactoe.Player{player1, player2}

	game := tictactoe.NewGame(board, players, ui)
	game.Run()
}
