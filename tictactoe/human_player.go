package tictactoe

import "strconv"

type Human struct {
  marker string
  move int
}

func (human *Human) SetMarker(marker string) {
  human.marker = marker
}

func (human *Human) Marker() string {
  return human.marker
}

func (human *Human) Move(board *Board) int {
  human.move,_ = strconv.Atoi(GetInput())
  return human.move
}
