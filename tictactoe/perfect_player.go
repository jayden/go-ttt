package tictactoe

const (
	initialDepth = 0
	depthLimit   = 6
	maxScore     = 10
)

type PerfectPlayer struct {
	marker string
	move   int
}

func NewPerfectPlayer() *PerfectPlayer {
	return new(PerfectPlayer)
}

func (player *PerfectPlayer) Marker() string {
	return player.marker
}

func (player *PerfectPlayer) SetMarker(marker string) {
	player.marker = marker
}

func (player *PerfectPlayer) Move(board *Board) int {
	player.minimax(board, player.marker, initialDepth)
	return player.move
}

func (player *PerfectPlayer) minimax(board *Board, marker string, depth int) int {
	if GameOver(board) || depth == depthLimit {
		return player.getEvaluatedScore(board, depth)
	}

	availableMoves := GetAvailableMoves(board)
	depth++
	scores := make(map[int]int, len(availableMoves))

	for _, move := range availableMoves {
		board.FillSpace(move, marker)
		nextTurn := getNextTurn(marker)
		scores[move] = player.minimax(board, nextTurn, depth)
		board.ClearSpace(move)
	}

	return player.getBestScoreForMarker(marker, scores)
}

func (player *PerfectPlayer) getBestScoreForMarker(marker string, scores map[int]int) int {
	var bestScore int
	if player.Marker() == marker {
		player.move, bestScore = getBestMove(scores, maxScore)
	} else {
		_, bestScore = getBestMove(scores, -maxScore)
	}
	return bestScore
}

func getBestMove(scores map[int]int, targetScore int) (int, int) {
	bestScore := targetScore * -1
	var bestMove int
	for move, score := range scores {
		if (targetScore > 0 && score > bestScore) || (targetScore < 0 && score < bestScore) {
			bestScore = score
			bestMove = move
		}
	}
	return bestMove, bestScore
}

func getNextTurn(marker string) string {
	if marker == x {
		return o
	}
	return x
}

func (player *PerfectPlayer) getEvaluatedScore(board *Board, depth int) int {
	winner, gameWon := GameWon(board)
	if gameWon {
		if winner == player.marker {
			return maxScore - depth
		} else {
			return depth - maxScore
		}
	}
	return 0
}
