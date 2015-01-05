package tictactoe

import (
  "fmt"
)

func GetInput() string {
  var input string
  fmt.Scan(&input)
  return input
}

func PrintMessage(message string) {
  fmt.Print(message)
}
