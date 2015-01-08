package tictactoe

const (
	blank            = " "
	defaultBoardSize = 9
)

type Board struct {
	spaces []string
}

func NewBoard() *Board {
	board := new(Board)
	initializeBoard(board)
	return board
}

func (board *Board) Spaces() []string {
	return board.spaces
}

func initializeBoard(board *Board) {
	board.spaces = make([]string, defaultBoardSize)
	for i := range board.spaces {
		board.spaces[i] = blank
	}
}

func (board *Board) FillSpace(position int, marker string) {
	board.spaces[position] = marker
}

func (board *Board) ClearSpace(position int) {
	board.FillSpace(position, blank)
}

func (board *Board) IsEmptySpace(position int) bool {
	return board.spaces[position] == blank
}
