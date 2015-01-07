package tictactoe

type MockConsoleUI struct {}

func (mockConsole *MockConsoleUI) AskForPlayerMove() int {
  return 0
}

func (mockConsole *MockConsoleUI) PrintBoard(board *Board) {

}

func (mockConsole *MockConsoleUI) PrintGameConclusion(board *Board) {

}

func (mockConsole *MockConsoleUI) PromptGameMenu() int {
  return 0
}
