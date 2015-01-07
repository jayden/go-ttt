package tictactoe

type UI interface {
  PromptGameMenu() int
  PrintBoard(*Board)
  PrintGameConclusion(*Board)
  AskForPlayerMove() int
}
