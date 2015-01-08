package tictactoe

import (
	"bytes"
	"io"
	"os"
)

func FillSpaces(board *Board, mark string, spaces ...int) {
	for _, space := range spaces {
		board.FillSpace(space, mark)
	}
}

func ClearBoard(board *Board) {
	for i, _ := range board.spaces {
		if !board.IsEmptySpace(i) {
			board.spaces[i] = blank
		}
	}
}

func writeMockInput(input string) {
	reader, writer, _ := os.Pipe()
	os.Stdin = reader
	writer.WriteString(input)
	writer.Close()
}

func writeMockOutput(message string, function func(string)) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	function(message)

	output := make(chan string)
	go func() {
		var buffer bytes.Buffer
		io.Copy(&buffer, r)
		output <- buffer.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-output

	return out
}
