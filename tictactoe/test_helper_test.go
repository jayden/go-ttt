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

func TestWriteMockOutput(t *testing.T) {
	printFunction := func(s string) { fmt.Print(s) }
	expected := "hello world!"
	actual := writeMockOutput(expected, printFunction)

	t.Log("helper method #writeMockOutput writes message to stdOut")
	assert.Equal(t, expected, actual)

}
