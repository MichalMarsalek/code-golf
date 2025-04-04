package hole

import "strings"

const reversiGridSize = 8

type (
	ReversiTile  int8
	ReversiBoard [reversiGridSize][reversiGridSize]ReversiTile
)

const (
	Empty ReversiTile = iota
	Black
	White
	PotentialSpot
)

func otherTeam(team ReversiTile) ReversiTile {
	if team == Black {
		return White
	} else if team == White {
		return Black
	} else {
		return team
	}
}

func inRange(x, y int) bool {
	return x >= 0 && y >= 0 && x < reversiGridSize && y < reversiGridSize
}

type ReversiSpot struct {
	Pos   [2]int
	Tiles [][2]int
}

func getPotentialSpots(team ReversiTile, grid ReversiBoard) []ReversiSpot {
	out := []ReversiSpot{}

	directions := [...][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
	}

	for i := range reversiGridSize {
		for j := range reversiGridSize {
			flippableTiles := [][2]int{}

			if grid[i][j] != Empty {
				continue
			}

			for _, direction := range directions {
				nextTileI := i + direction[0]
				nextTileJ := j + direction[1]

				if !inRange(nextTileJ, nextTileI) {
					continue
				}

				if grid[nextTileI][nextTileJ] != otherTeam(team) {
					continue
				}

				flippableTilesInDirection := [][2]int{}

				for inRange(nextTileJ, nextTileI) && grid[nextTileI][nextTileJ] == otherTeam(team) {
					flippableTilesInDirection = append(flippableTilesInDirection, [2]int{nextTileI, nextTileJ})

					nextTileI += direction[0]
					nextTileJ += direction[1]
				}

				if inRange(nextTileJ, nextTileI) && grid[nextTileI][nextTileJ] == team {
					flippableTiles = append(flippableTiles, flippableTilesInDirection...)
				}
			}

			if len(flippableTiles) > 0 {
				out = append(out, ReversiSpot{
					Pos:   [2]int{i, j},
					Tiles: flippableTiles,
				})
			}
		}
	}

	return out
}

func getReversiInitialState() (out ReversiBoard) {
	out[3][4] = Black
	out[4][3] = Black
	out[3][3] = White
	out[4][4] = White
	return out
}

func flipReversiBoard(board ReversiBoard, bottomLeft, bottomRight bool) ReversiBoard {
	outBoard := getReversiInitialState()

	for x := range reversiGridSize {
		for y := range reversiGridSize {
			newX := x
			newY := y

			if bottomLeft {
				newX = y
				newY = x
			}

			if bottomRight {
				newX = reversiGridSize - 1 - newX
				newY = reversiGridSize - 1 - newY
			}

			outBoard[newX][newY] = board[x][y]
		}
	}

	return outBoard
}

func genReversiBoard(steps int) ReversiBoard {
	for {
		out := getReversiInitialState()
		valid := true

		for i := range steps {
			team := []ReversiTile{Black, White}[i%2]

			spots := getPotentialSpots(team, out)
			if len(spots) == 0 {
				valid = false
				break
			}

			spot := randChoice(spots)

			out[spot.Pos[0]][spot.Pos[1]] = team
			for _, reversedSpot := range spot.Tiles {
				out[reversedSpot[0]][reversedSpot[1]] = team
			}

		}

		if valid {
			return out
		}
	}
}

func drawReversiBoard(board ReversiBoard) string {
	var sb strings.Builder

	for i := range reversiGridSize {
		if i != 0 {
			sb.WriteByte('\n')
		}

		for j := range reversiGridSize {
			if board[i][j] == Empty {
				sb.WriteByte('.')
			} else if board[i][j] == Black {
				sb.WriteByte('X')
			} else if board[i][j] == White {
				sb.WriteByte('O')
			} else if board[i][j] == PotentialSpot {
				sb.WriteByte('!')
			}
		}
	}

	return sb.String()
}

func highlightCorrectAnswersReversiBoard(board ReversiBoard) ReversiBoard {
	for _, spot := range getPotentialSpots(White, board) {
		board[spot.Pos[0]][spot.Pos[1]] = PotentialSpot
	}
	return board
}

func getStaticTestCases() (boards []ReversiBoard) {
	for _, test := range fixedTests("reversi") {
		board := getReversiInitialState()

		// Copy test.in to board.
		for i, line := range strings.Split(test.in, "\n") {
			for j, c := range line {
				if c == 'X' {
					board[i][j] = Black
				} else if c == 'O' {
					board[i][j] = White
				}
			}
		}

		boards = append(
			boards,
			flipReversiBoard(board, true, false),
			flipReversiBoard(board, false, true),
			flipReversiBoard(board, true, true),
		)
	}

	return shuffle(boards)
}

func getAutogeneratedTestCases() []ReversiBoard {
	const runs = 20

	out := make([]ReversiBoard, runs)
	for i := range runs {
		grid := genReversiBoard((i+1)/2*2 + 1)
		out[i] = grid
	}

	return out
}

var _ = answerFunc("reversi", func() []Answer {
	runs := make([][]test, 2)

	boards := [][]ReversiBoard{
		getAutogeneratedTestCases(),
		getStaticTestCases(),
	}

	for i, board := range boards {
		runs[i] = make([]test, len(board))

		for j, grid := range board {
			runs[i][j].in = drawReversiBoard(grid)
			runs[i][j].out = drawReversiBoard(highlightCorrectAnswersReversiBoard(grid))
		}
	}

	return outputTestsWithSep("\n\n", runs...)
})
