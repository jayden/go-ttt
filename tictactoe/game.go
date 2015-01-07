package tictactoe

import (
  "regexp"
  "strconv"
)

type Game struct {
  board *Board
  players []Player
  ui UI
}

func NewGame(board *Board, players []Player, ui UI) *Game {
  game := new(Game)
  game.board = board
  game.players = players
  game.ui = ui
  return game
}

func (game *Game) Board() *Board {
  return game.board
}

func (game *Game) PutMove(position int, marker string) {
  game.board.FillSpace(position, marker)
}

func (game *Game) IsValidMove(input string) bool {
  isNumber,_ := regexp.MatchString("^[0-8]$", input)
  if isNumber {
    position,_ := strconv.Atoi(input)
    return game.isValidPosition(position)
  }
  return false
}

func (game *Game) isValidPosition(position int) bool {
  board := game.board
  valid := position >= 0 && position < len(board.spaces)
  return valid && board.IsEmptySpace(position)
}
