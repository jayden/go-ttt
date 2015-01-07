package tictactoe

type Player interface {
  Marker() string
  SetMarker(string)
  Move(*Board) int
}
