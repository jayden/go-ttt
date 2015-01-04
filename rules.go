package tictactoe

func (board *Board) GameDraw() bool {
  if containsBlankSpace(board) {
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

func (board *Board) GameWon() bool {
  if markersInSet(board, []int{0,1,2}, "X") || markersInSet(board, []int{3,4,5}, "O") {
    return true
  }
  return false
}

func markersInSet(board *Board, combo []int, marker string) bool {
  result := make([]string, len(combo))
  for i,position := range combo {
    result[i] = board.spaces[position]
  }

  unique := true
  for _,mark := range result {
    unique = unique && mark == marker
  }
  return unique
}
