package tictactoe

import "strconv"

type Human struct {
  marker string
  move int
}

func (human *Human) SetMarker(marker string) {
  human.marker = marker
}

func (human *Human) GetMarker() string {
  return human.marker
}

func (human *Human) GetMove(board *Board) {
  human.move,_ = strconv.Atoi(GetInput())
  board.Fill(human.move, human.marker)
}
