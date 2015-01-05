package tictactoe

import (
  "bytes"
  "github.com/stretchr/testify/assert"
  "io"
  "os"
  "testing"
)

func TestGetInput(t *testing.T) {
  t.Log("gets input from console")
  mockInput := "1"
  writeMockInput(mockInput)
  assert.Equal(t, mockInput, GetInput())
}

func writeMockInput(input string) {
  reader,writer,_ := os.Pipe()
  os.Stdin = reader
  writer.WriteString(input)
  writer.Close()
}

func TestPrintsOutput(t *testing.T) {
  t.Log("shows output to console")
  old := os.Stdout
  r, w, _ := os.Pipe()
  os.Stdout = w
  message := "hello world!"
  PrintMessage(message)

  outC := make(chan string)
  go func() {
    var buffer bytes.Buffer
    io.Copy(&buffer, r)
    outC <- buffer.String()
  }()

  w.Close()
  os.Stdout = old
  out := <-outC

  assert.Equal(t, message, out)
}

func TestPrintsBoard(t *testing.T) {
  t.Log("Prints blank board")
  expected := "   |   |   \n-----------\n   |   |   \n-----------\n   |   |   "
  assert.Equal(t, expected, BoardToString(MakeBoard()))

  t.Log("Prints board with markers")
  expected = " X | X | O \n-----------\n O | O | X \n-----------\n X | X | O "
  assert.Equal(t, expected, BoardToString(tieBoard()))
}
