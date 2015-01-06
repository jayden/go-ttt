package tictactoe

const (
  blank = " "
  defaultBoardSize = 9
)

type Board struct {
  spaces []string
}

func MakeBoard() *Board {
  board := new(Board)
  initializeBoard(board)
  return board
}

func initializeBoard(board *Board) {
  board.spaces = make([]string, defaultBoardSize)
  for i:= range board.spaces {
    board.spaces[i] = blank
  }
}

func (board *Board) FillSpace(position int, marker string) {
  if validPosition(board, position) {
    board.spaces[position] = marker
  }
}

func (board *Board) ClearSpace(position int) {
  board.FillSpace(position, blank)
}

func validPosition(board *Board, position int) bool {
  return position >= 0 && position < len(board.spaces)
}
