package tictactoe

func (board *Board) GameDraw() bool {
  _,gameWon := board.GameWon()
  if containsBlankSpace(board) || gameWon {
    return false
  }
  return true
}

func containsBlankSpace(board *Board) bool {
  for _, marker := range board.spaces {
    if marker == blank {
      return true
    }
  }
  return false
}

func (board *Board) GameWon() (string, bool) {
  for _,marker := range []string{"X", "O"} {
    for _,combo := range winningCombos() {
     set := markersInCombo(board, combo)
     if comboHasWinner(set, marker) { return marker, true }
   }
  }
  return "", false
}

func markersInCombo(board *Board, combo []int) []string {
  result := make([]string, len(combo))
  for i,position := range combo {
    result[i] = board.spaces[position]
  }
  return result
}

func comboHasWinner(combo []string, marker string) bool {
  unique := true
  for _,mark := range combo {
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

func (board *Board) GetAvailableMoves() []int {
  result := []int{}
  if board.GameOver() {
    return result
  } else {
    for i,space := range board.spaces {
      if space == blank { result = append(result, i) }
    }
  }
    return result
}

func (board *Board) GameOver() bool {
  _, gameWon := board.GameWon()
  return board.GameDraw() || gameWon
}
