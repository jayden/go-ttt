package tictactoe

type MockConsoleUI struct {
	move           int
	selection      int
	calledFunction string
}

func (mockConsole *MockConsoleUI) MakeFakeGameSelection(selection int) {
	mockConsole.selection = selection
}

func (mockConsole *MockConsoleUI) PromptGameMenu() int {
	mockConsole.calledFunction = "PromptGameMenu"
	return mockConsole.selection
}

func (mockConsole *MockConsoleUI) MakeFakePlayerMove(move int) {
	mockConsole.move = move
}

func (mockConsole *MockConsoleUI) AskForPlayerMove(board *Board) int {
	mockConsole.calledFunction = "AskForPlayerMove"
	return mockConsole.move
}

func (mockConsole *MockConsoleUI) PrintBoard(board *Board) {
	mockConsole.calledFunction = "PrintBoard"
}

func (mockConsole *MockConsoleUI) PrintGameConclusion(board *Board) {
	mockConsole.calledFunction = "PrintGameConclusion"
}

func (mockConsole *MockConsoleUI) Called(function string) bool {
	if mockConsole.calledFunction == function {
		return true
	}
	return false
}

func (mockConsole *MockConsoleUI) ResetCall() {
	mockConsole.calledFunction = ""
}
