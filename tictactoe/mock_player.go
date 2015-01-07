package tictactoe

type MockPlayer struct {
  marker string
  move int
}

func (mockPlayer *MockPlayer) Marker() string {
  return mockPlayer.marker
}

func (mockPlayer *MockPlayer) SetMarker(marker string) {
  mockPlayer.marker = marker
}

func (mockPlayer *MockPlayer) MakeFakeMove(move int) {
  mockPlayer.move = move
}

func (mockPlayer *MockPlayer) Move(board *Board) int {
  return mockPlayer.move
}
