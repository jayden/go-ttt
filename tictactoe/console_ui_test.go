package tictactoe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsoleIsAUI(t *testing.T) {
	console := new(ConsoleUI)
	var ui UI = console

	t.Log("Console UI is a kind of UI")
	assert.Equal(t, console, ui)
}

func TestGetInput(t *testing.T) {
	mockedInput := "1"
	writeMockInput(mockedInput)
	input := GetInput()

	t.Log("gets input from console")
	assert.Equal(t, mockedInput, input)
}

func TestPrintsOutput(t *testing.T) {
	message := "hello world!"
	printFunction := func(s string) { PrintMessage(s) }
	mockedOutput := writeMockOutput(message, printFunction)

	t.Log("shows output to console")
	assert.Equal(t, message, mockedOutput)
}

func TestPrintsBoard(t *testing.T) {
	t.Log("Prints blank board")
	expected := "   |   |   \n-----------\n   |   |   \n-----------\n   |   |   "
	assert.Equal(t, expected, BoardToString(MakeBoard()))

	t.Log("Prints board with markers")
	expected = " X | X | O \n-----------\n O | O | X \n-----------\n X | X | O "
	assert.Equal(t, expected, BoardToString(tieBoard()))
}

func TestValidatesInput(t *testing.T) {
	t.Log("returns false if it's not an integer")
	assert.False(t, isValidInput("m"))

	t.Log("returns true if it's a positive integer")
	assert.True(t, isValidInput("1"))

	t.Log("returns false if it's a negative integer")
	assert.False(t, isValidInput("-1"))
}
