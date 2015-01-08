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
	game.setPlayerMarkers()
	return game
}

func (game *Game) setPlayerMarkers() {
	game.players[0].SetMarker("X")
	game.players[1].SetMarker("O")
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
	currentPlayer, nextPlayer := game.players[0], game.players[1]
	for !GameOver(game.board) {
		game.ui.PrintBoard(game.board)
		move := currentPlayer.Move(game.board)
		game.PutMove(move, currentPlayer.Marker())
		currentPlayer, nextPlayer = nextPlayer, currentPlayer
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
