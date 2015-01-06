package tictactoe

type Player interface {
  GetMarker() string
  SetMarker(string)
  GetMove(*Board) int
}
