package tictactoe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockConsoleIsAUI(t *testing.T) {
	mockConsole := new(MockConsoleUI)
	var ui UI = mockConsole

	t.Log("Mock Console UI is a kind of UI")
	assert.Equal(t, mockConsole, ui)
}

func TestMockConsoleMakesFakeGameSelection(t *testing.T) {
	mockConsole := new(MockConsoleUI)
	mockConsole.MakeFakeGameSelection(1)

	t.Log("Mock Console makes fake game selection")
	assert.Equal(t, 1, mockConsole.PromptGameMenu())
}

func TestMockConsoleMakesFakePlayerMove(t *testing.T) {
	mockConsole := new(MockConsoleUI)
	board := NewBoard()
	mockConsole.MakeFakePlayerMove(1)

	t.Log("Mock Console makes fake player move")
	assert.Equal(t, 1, mockConsole.AskForPlayerMove(board))
}

func TestMockConsoleKeepsTrackOfCalledFunctions(t *testing.T) {
	mockConsole := new(MockConsoleUI)
	mockConsole.PromptGameMenu()

	t.Log("Mock Console records that #PromptGameMenu has been called")
	assert.True(t, mockConsole.Called("PromptGameMenu"))

	board := NewBoard()
	mockConsole.AskForPlayerMove(board)

	t.Log("Mock Console records that #AskForPlayerMove has been called")
	assert.True(t, mockConsole.Called("AskForPlayerMove"))

	mockConsole.PrintBoard(board)

	t.Log("Mock Console records that #PrintBoard has been called")
	assert.True(t, mockConsole.Called("PrintBoard"))

	mockConsole.PrintGameConclusion(board)

	t.Log("Mock Console records that #PrintGameConclusion has been called")
	assert.True(t, mockConsole.Called("PrintGameConclusion"))
}

func TestMockConsoleResetsTrackedCall(t *testing.T) {
	mockConsole := new(MockConsoleUI)
	mockConsole.PromptGameMenu()

	t.Log("Mock Console resets called function")
	assert.True(t, mockConsole.Called("PromptGameMenu"))
	mockConsole.ResetCall()
	assert.False(t, mockConsole.Called("PromptGameMenu"))
}
