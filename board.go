package tictactoe

type Board struct {
  spaces []string
}

func MakeBoard() *Board {
  board := new(Board)
  board.spaces = make([]string, 9)
  for i:= range board.spaces {
    board.spaces[i] = " "
  }
  return board
}
