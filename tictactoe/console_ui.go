package tictactoe

import (
  "bytes"
  "fmt"
  "math"
)

const (
  horizontalLine = "\n-----------\n"
  spaceSeparator = "|"
)

func GetInput() string {
  var input string
  fmt.Scan(&input)
  return input
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
  if isLastColumn(index){ return writeHorizontalLine(index) }
  return spaceSeparator
}

func isLastColumn(index int) bool {
  i := float64(index+1)
  return math.Mod(i, 3) == 0
}

func writeHorizontalLine(index int) string {
  if index+1 < 9 { return horizontalLine }
  return ""
}
