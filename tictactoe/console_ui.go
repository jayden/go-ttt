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
)

type ConsoleUI struct{}

func (console ConsoleUI) PromptGameMenu() int {
	PrintMessage(welcomeMessage)
	menu := "1. Human vs. Computer\n" +
		"2. Computer vs. Human\n" +
		"3. Exit Game\n\n" +
		"Please select a game: "
	PrintMessage(menu)
	selection := convertInputToInt(GetInput())
	return selection
}

func (console ConsoleUI) PrintBoard(board *Board) {
	PrintMessage(BoardToString(board))
}

func (console ConsoleUI) PrintGameConclusion(board *Board) {
	if board.GameDraw() {
		PrintMessage(drawMessage)
	} else {
		winner, _ := board.GameWon()
		winMessage := winner + " wins!\n"
		PrintMessage(winMessage)
	}
}

func convertInputToInt(input string) int {
	result, _ := strconv.Atoi(input)
	return result
}

func (console ConsoleUI) AskForPlayerMove() int {
	for {
		PrintMessage(askForMoveMessage)
		input := GetInput()
		if !isValidInput(input) {
			PrintMessage(invalidMoveMessage)
			continue
		}
		return convertInputToInt(input)
	}
}

func isValidInput(input string) bool {
	isMatch, _ := regexp.MatchString("^[0-9]$", input)
	return isMatch
}

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
	for i, mark := range board.spaces {
		buffer.WriteString(" " + mark + " ")
		buffer.WriteString(writeSeparator(i))
	}
	return buffer.String()
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
