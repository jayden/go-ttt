package tictactoe

const (
	x = "x"
	o = "o"
)

func GameDraw(board *Board) bool {
	_, gameWon := GameWon(board)
	if containsBlankSpace(board) || gameWon {
		return false
	}
	return true
}

func containsBlankSpace(board *Board) bool {
	for _, marker := range board.Spaces() {
		if marker == blank {
			return true
		}
	}
	return false
}

func GameWon(board *Board) (string, bool) {
	for _, marker := range []string{x, o} {
		for _, combo := range winningCombos() {
			set := markersInCombo(board, combo)
			if comboHasWinner(set, marker) {
				return marker, true
			}
		}
	}
	return "", false
}

func markersInCombo(board *Board, combo []int) []string {
	result := make([]string, len(combo))
	for i, position := range combo {
		result[i] = board.Spaces()[position]
	}
	return result
}

func comboHasWinner(combo []string, marker string) bool {
	unique := true
	for _, mark := range combo {
		unique = unique && mark == marker
	}
	return unique
}

func winningCombos() [][]int {
	return [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}
}

func GetAvailableMoves(board *Board) []int {
	result := []int{}
	if GameOver(board) {
		return result
	} else {
		for i, space := range board.Spaces() {
			if space == blank {
				result = append(result, i)
			}
		}
	}
	return result
}

func GameOver(board *Board) bool {
	_, gameWon := GameWon(board)
	return GameDraw(board) || gameWon
}
