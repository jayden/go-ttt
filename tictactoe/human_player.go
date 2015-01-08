package tictactoe

type Human struct {
	marker  string
	console UI
}

func NewHumanPlayer() *Human {
	human := new(Human)
	human.console = new(ConsoleUI)
	return human
}

func (human *Human) SetMarker(marker string) {
	human.marker = marker
}

func (human *Human) Marker() string {
	return human.marker
}

func (human *Human) Move(board *Board) int {
	return human.console.AskForPlayerMove(board)
}
