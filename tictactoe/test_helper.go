package tictactoe

func FillSpaces(board *Board, mark string, spaces ...int) {
  for _, space := range spaces {
    board.FillSpace(space, mark)
  }
}
