package tictactoe

import (
  "fmt"
  "sort"
)

type PerfectPlayer struct {
  marker string
  move int
  bestMove int
}

func (player *PerfectPlayer) GetMarker() string {
  return player.marker
}

func (player *PerfectPlayer) SetMarker(marker string) {
  player.marker = marker
}

func (player *PerfectPlayer) GetOpponentMarker() string {
  if player.marker == "X" { return "O" }
  return "X"
}

func (player *PerfectPlayer) GetMove(board *Board) int {
  player.minimax(board, player.marker, 0)
  return player.bestMove
}

func (player *PerfectPlayer) minimax(board *Board, marker string, depth int) int {
  if board.GameOver() {
    return player.getEvaluatedScore(board, depth)
  }

  availableMoves := board.GetAvailableMoves()
  fmt.Println(availableMoves)
  depth++
  scores, moves := []int{}, []int{}
  scoresMap := make(map[int]int, len(board.spaces))

  for _,move := range availableMoves {
    board.Fill(move, marker)
    nextTurn := getNextTurn(marker)
    recurScore := player.minimax(board, nextTurn, depth)
    scores = append(scores, recurScore)
    moves = append(moves, move)
    scoresMap[move] = recurScore
    board.ClearSpace(move)
    fmt.Println("SCORES: ", scores)
    fmt.Println("MOVES: ", moves)
  }

  fmt.Println("SCORES: ", scores)
  fmt.Println("MOVES: ", moves)
  fmt.Println("MAP: ", scoresMap)

  minScoreIndex, maxScoreIndex := getIndicesOfBestScores(scores)
  if player.GetMarker() == marker {
    player.bestMove =  moves[maxScoreIndex]
    return scores[maxScoreIndex]
  } else {
    player.bestMove =  moves[minScoreIndex]
    return scores[minScoreIndex]
  }
}

func getNextTurn(marker string) string {
  if marker == "X" { return "O" }
  return "X"
}

func getIndicesOfBestScores(scores []int) (int, int) {
  sortedScores := make([]int, len(scores))
  copy(sortedScores[:], scores)
  sort.Ints(sortedScores)
  minScore, maxScore := 0, len(sortedScores) - 1
  var minScoreIndex, maxScoreIndex int

  for i, score := range scores {
    if score == sortedScores[minScore] {
      minScoreIndex = i
    } else if score == sortedScores[maxScore] {
      maxScoreIndex = i
    }
  }
  fmt.Println("MINSCOREINDEX: ", minScoreIndex)
  fmt.Println("MAXSCOREINDEX: ", maxScoreIndex)
  return minScoreIndex, maxScoreIndex
}

func (player *PerfectPlayer) getEvaluatedScore(board *Board, depth int) int {
  baseScore := len(board.spaces) + 1
  fmt.Println("DEPTH FOR BOARD: ", board.spaces, depth)
  winner, gameWon := board.GameWon()
  if gameWon {
    if winner == player.marker {
      fmt.Println("SCORE IF WINNER = COMP: ", baseScore - depth)
      return baseScore - depth
    } else {
      fmt.Println("SCORE IF WINNER = HUMAN: ", depth - baseScore)
      return depth - baseScore
    }
  } else {
    return 0
  }
}