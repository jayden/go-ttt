package tictactoe

type Game struct {
	board   *Board
	players []Player
	ui      UI
}

const (
	humanFirst    = 1
	computerFirst = 2
	exit          = 3
)

func NewGame(board *Board, players []Player, ui UI) *Game {
	game := new(Game)
	game.board = board
	game.players = players
	game.ui = ui
	return game
}

func (game *Game) Board() *Board {
	return game.board
}

func (game *Game) PutMove(position int, marker string) {
	game.board.FillSpace(position, marker)
}

func (game *Game) Run() {
	selection := game.ui.PromptGameMenu()
	switch selection {
	case humanFirst:
		break
	case computerFirst:
		game.changePlayerOrder()
	case exit:
		return
	}
	game.runGame()
}

func (game *Game) runGame() {
	current, next := game.players[0], game.players[1]
	for !game.board.GameOver() {
		game.ui.PrintBoard(game.board)
		move := current.Move(game.board)
		game.PutMove(move, current.Marker())
		current, next = next, current
	}
	game.ui.PrintBoard(game.board)
	game.ui.PrintGameConclusion(game.board)
}

func (game *Game) changePlayerOrder() {
	firstMarker := game.players[0].Marker()
	secondMarker := game.players[1].Marker()
	game.players[0], game.players[1] = game.players[1], game.players[0]
	game.players[0].SetMarker(firstMarker)
	game.players[1].SetMarker(secondMarker)
}
