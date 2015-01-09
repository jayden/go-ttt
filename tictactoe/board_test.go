package tictactoe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeBlankBoard(t *testing.T) {
	actual := NewBoard().spaces
	expected := []string{blank, blank, blank, blank, blank, blank, blank, blank, blank}

	t.Log("can create a blank board")
	assert.Equal(t, expected, actual)
}

func TestSpaces(t *testing.T) {
	board := NewBoard()
	board.spaces[0] = x

	t.Log("return a copy of the board spaces")
	assert.Equal(t, board.spaces, board.Spaces())
}

func TestFillBoard(t *testing.T) {
	board := NewBoard()
	board.FillSpace(4, x)

	t.Log("can fill a board with a marker")
	assert.Equal(t, board.spaces[4], x)
}

func TestClearSpace(t *testing.T) {
	board := NewBoard()
	board.FillSpace(4, x)

	t.Log("can clear a space on the board")
	assert.Equal(t, board.spaces[4], x)
	board.ClearSpace(4)
	assert.Equal(t, board.spaces[4], blank)
}
