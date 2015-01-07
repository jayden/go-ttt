package tictactoe

import (
  "bytes"
  "fmt"
  "math"
  "strconv"
)

const (
  askForMoveMessage = "Please select a move: "
  drawMessage = "It's a draw!\n"
  horizontalLine = "\n-----------\n"
  spaceSeparator = "|"
  welcomeMessage = "Let's Play Tic-Tac-Toe!\n"
)

type ConsoleUI struct {}

func (console ConsoleUI) PromptGameMenu() int {
  PrintMessage(welcomeMessage)
  menu := "1. Human vs. Computer\n" +
          "2. Computer vs. Human\n" +
          "3. Exit Game\n\n" +
          "Please select a game: "
  PrintMessage(menu)
  selection := GetInput()
  return selection
}

func (console ConsoleUI) PrintBoard(board *Board) {
  PrintMessage(BoardToString(board))
}

func (console ConsoleUI) PrintGameConclusion(board *Board) {
  console.PrintBoard(board)

  if board.GameDraw() {
    PrintMessage(drawMessage)
  } else {
    winner,_ := board.GameWon()
    winMessage := winner + " wins!\n"
    PrintMessage(winMessage)
  }
}

func (console ConsoleUI) AskForPlayerMove() int {
  PrintMessage(askForMoveMessage)
  return GetInput()
}

func GetInput() int {
  var input string
  fmt.Scan(&input)
  result,_ := strconv.Atoi(input)
  return result
}

func PrintMessage(message string) {
  fmt.Print(message)
}

func BoardToString(board *Board) string {
  var buffer bytes.Buffer
  for i,mark := range board.spaces {
    buffer.WriteString(" " + mark + " ")
    buffer.WriteString(writeSeparator(i))
  }
  return buffer.String()
}

func writeSeparator(index int) string {
  if isLastSpaceInRow(index){ return writeHorizontalLine(index) }
  return spaceSeparator
}

func isLastSpaceInRow(index int) bool {
  i := float64(index+1)
  return math.Mod(i, 3) == 0
}

func writeHorizontalLine(index int) string {
  if index+1 < 9 { return horizontalLine }
  return ""
}
