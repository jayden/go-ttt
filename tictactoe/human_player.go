package tictactoe

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
  human.move = GetInput()
  return human.move
}
