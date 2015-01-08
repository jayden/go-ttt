package tictactoe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFillSpaces(t *testing.T) {
	t.Log("helper method #FillSpaces fills multiple spaces at once")
	board := MakeBoard()
	FillSpaces(board, "X", 0, 1, 2)
	expected := []string{"X", "X", "X", blank, blank, blank, blank, blank, blank}
	assert.Equal(t, expected, board.spaces)
}

func TestClearBoard(t *testing.T) {
	board := MakeBoard()
	FillSpaces(board, "X", 0, 1, 2)
	expectedBefore := []string{"X", "X", "X", blank, blank, blank, blank, blank, blank}
	expectedAfter := []string{blank, blank, blank, blank, blank, blank, blank, blank, blank}

	t.Log("helper method #ClearBoard clears the entire board")
	assert.Equal(t, expectedBefore, board.spaces)
	ClearBoard(board)
	assert.Equal(t, expectedAfter, board.spaces)
}

func TestWriteMockOutput(t *testing.T) {
	printFunction := func(s string) { fmt.Print(s) }
	expected := "hello world!"
	actual := writeMockOutput(expected, printFunction)

	t.Log("helper method #writeMockOutput writes message to stdOut")
	assert.Equal(t, expected, actual)

}
