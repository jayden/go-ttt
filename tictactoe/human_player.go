package tictactoe

type Human struct {
	marker  string
	move    int
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
	human.move = human.console.AskForPlayerMove()
	return human.move
}
