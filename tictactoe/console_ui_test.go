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

func TestAskForPlayerMove(t *testing.T) {
	console := new(ConsoleUI)
	mockedInput := "1"
	writeMockInput(mockedInput)
	board := MakeBoard()

	move := console.AskForPlayerMove(board)

	t.Log("should return move successfully on valid input")
	assert.Equal(t, 1, move)
}

func TestPromptGameMenu(t *testing.T) {
	console := new(ConsoleUI)
	mockedInput := "1"
	writeMockInput(mockedInput)
	printMenu := func(s string) { PrintMessage(gameMenu) }
	printWelcome := func(s string) { PrintMessage(welcomeMessage) }
	outputMenu := writeMockOutput(gameMenu, printMenu)
	outputWelcome := writeMockOutput(welcomeMessage, printWelcome)

	selection := console.PromptGameMenu()

	t.Log("it should print welcome message")
	assert.Equal(t, welcomeMessage, outputWelcome)

	t.Log("it should print menu")
	assert.Equal(t, gameMenu, outputMenu)

	t.Log("it returns the user selection on valid input")
	assert.Equal(t, 1, selection)
}

func TestChecksIfInputIsNumber(t *testing.T) {
	t.Log("returns false if it's not an number")
	assert.False(t, isNumber("m"))

	t.Log("returns true if it's a positive number")
	assert.True(t, isNumber("1"))

	t.Log("returns false if it's a negative number")
	assert.False(t, isNumber("-1"))
}

func TestChecksIfInputIsValid(t *testing.T) {
	t.Log("returns true if input is within list of valid inputs")
	assert.True(t, isValidInput(1, []int{1, 2, 3}))

	t.Log("returns false if input is not within list")
	assert.False(t, isValidInput(2, []int{1, 3}))
}
