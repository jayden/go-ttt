package tictactoe

import (
	"bytes"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

const (
	askForMoveMessage  = "Please select a move: "
	drawMessage        = "It's a draw!\n"
	horizontalLine     = "\n-----------\n"
	invalidMoveMessage = "Invalid move!\n"
	spaceSeparator     = "|"
	welcomeMessage     = "Let's Play Tic-Tac-Toe!\n"
	gameMenu           = "1. Human vs. Computer\n" +
		"2. Computer vs. Human\n" +
		"3. Exit Game\n\n" +
		"Please select a game: "
)

type ConsoleUI struct{}

func NewConsoleUI() *ConsoleUI {
	return new(ConsoleUI)
}

func (console ConsoleUI) PromptGameMenu() int {
	PrintMessage(welcomeMessage)
	validSelections := []int{1, 2, 3}

	for {
		PrintMessage(gameMenu)
		input := GetInput()
		if isNumber(input) {
			selection := convertInputToInt(input)
			if isValidInput(selection, validSelections) {
				return selection
			}
		}
		continue
	}
}

func PrintMessage(message string) {
	fmt.Print(message)
}

func GetInput() string {
	var input string
	fmt.Scan(&input)
	return input
}

func (console ConsoleUI) PrintBoard(board *Board) {
	PrintMessage("\n" + BoardToString(board) + "\n\n")
}

func (console ConsoleUI) PrintGameConclusion(board *Board) {
	if GameDraw(board) {
		PrintMessage(drawMessage)
	} else {
		winner, _ := GameWon(board)
		winMessage := winner + " wins!\n"
		PrintMessage(winMessage)
	}
}

func convertInputToInt(input string) int {
	result, _ := strconv.Atoi(input)
	return result
}

func (console ConsoleUI) AskForPlayerMove(board *Board) int {
	for {
		PrintMessage(askForMoveMessage)
		input := GetInput()
		if isNumber(input) {
			move := convertInputToInt(input)
			if isValidInput(move, GetAvailableMoves(board)) {
				return move
			}
		}
		PrintMessage(invalidMoveMessage)
		continue
	}
}

func isValidInput(input int, validInputs []int) bool {
	for _, value := range validInputs {
		if value == input {
			return true
		}
	}
	return false
}

func isNumber(input string) bool {
	isMatch, _ := regexp.MatchString("^[0-9]$", input)
	return isMatch
}

func BoardToString(board *Board) string {
	var buffer bytes.Buffer
	for i, mark := range board.spaces {
		if mark == blank {
			mark = blankSpaceToIndex(i)
		}
		buffer.WriteString(" " + mark + " ")
		buffer.WriteString(writeSeparator(i))
	}
	return buffer.String()
}

func blankSpaceToIndex(index int) string {
	return strconv.Itoa(index)
}

func writeSeparator(index int) string {
	if isLastSpaceInRow(index) {
		return writeHorizontalLine(index)
	}
	return spaceSeparator
}

func isLastSpaceInRow(index int) bool {
	i := float64(index + 1)
	return math.Mod(i, 3) == 0
}

func writeHorizontalLine(index int) string {
	if index+1 < 9 {
		return horizontalLine
	}
	return ""
}
