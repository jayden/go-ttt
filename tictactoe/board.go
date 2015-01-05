package tictactoe

type Board struct {
  spaces []string
}

const blank = " "

func MakeBoard() *Board {
  board := new(Board)
  initializeBoard(board)
  return board
}

func initializeBoard(board *Board) {
  board.spaces = make([]string, 9)
  for i:= range board.spaces {
    board.spaces[i] = blank
  }
}

func (board *Board) Fill(position int, marker string) {
  if validFillPosition(board, position) {
    board.spaces[position] = marker
  }
}

func (board *Board) ClearSpace(position int) {
  if validFillPosition(board, position) {
    board.spaces[position] = blank
  }
}

func validFillPosition(board *Board, position int) bool {
  return position >= 0 && position < len(board.spaces)
}

