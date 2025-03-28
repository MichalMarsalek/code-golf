package hole

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
)

const (
	cols int = 7
	rows int = 6
)

func dropDisc(grid [][]int, col, player int) bool {
	for i := rows - 1; i >= 0; i-- {
		if grid[i][col] == 0 {
			grid[i][col] = player

			return true
		}
	}

	return false
}

func checkDirection(grid [][]int, row, col, dr, dc, player int) bool {
	count := 0

	for i := range 4 {
		if nr, nc := row+dr*i, col+dc*i; nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == player {
			count++
		} else {
			break
		}
	}

	return count == 4
}

func checkWinner(grid [][]int, player int) int {
	for row := range rows {
		for col := range cols {
			if grid[row][col] == player {
				for _, pts := range [][]int{
					{1, 0}, {0, 1}, {1, 1}, {1, -1},
				} {
					if checkDirection(grid, row, col, pts[0], pts[1], player) {
						return player
					}
				}
			}
		}
	}

	return 0
}

func emulPlay() ([]int, string) {
	grid, player := make([][]int, rows), 1

	for i := range grid {
		grid[i] = make([]int, cols)
	}

	var moves []int

	for {
		if col := rand.IntN(cols); dropDisc(grid, col, player) {
			moves = append(moves, col)

			if winner := checkWinner(grid, player); winner != 0 {
				return moves, []string{"Red", "Yellow"}[player-1]
			}

			// If there are no empty cells on the top row, return "Draw".
			if !slices.Contains(grid[0], 0) {
				return moves, "Draw"
			}

			player = 3 - player
		}

		// Randomly terminate a game in progress
		if len(moves) >= 7 && rand.IntN(99) == 0 {
			return moves, "-"
		}
	}
}

var _ = answerFunc("connect-four", func() []Answer {
	draws, sweats := shuffle([]string{
		"0 0 2 4 5 2 5 0 2 1 6 3 3 3 6 4 4 5 2 5 5 3 1 3 3 0 5 6 0 0 4 2 1 4 2 4 6 1 6 6 1 1",
		"0 0 3 1 2 6 2 0 6 0 0 6 3 2 6 0 5 4 3 6 2 2 1 6 5 2 5 5 5 4 1 3 3 1 3 5 4 4 4 4 1 1",
		"0 0 3 3 4 6 6 6 0 4 0 5 6 6 3 1 4 0 0 3 2 2 1 1 2 2 5 5 6 2 2 4 1 4 3 3 5 4 5 1 5 1",
		"0 0 3 6 5 6 6 6 4 1 1 1 3 3 5 0 0 1 6 0 0 2 4 3 1 2 2 1 6 5 2 3 3 4 4 4 2 2 5 5 5 4",
		"0 0 4 4 3 2 4 3 0 6 1 6 4 1 0 3 2 0 5 0 5 6 6 1 6 4 2 4 2 1 1 1 6 5 3 2 2 3 3 5 5 5",
		"0 0 5 3 5 3 6 6 6 1 3 4 2 0 2 1 4 3 2 4 3 6 3 2 4 2 2 5 4 4 0 1 1 5 0 1 1 5 5 0 6 6",
		"0 0 6 0 2 0 6 1 1 4 6 5 0 0 1 5 2 5 5 6 3 6 6 2 2 2 1 3 2 4 4 4 4 1 3 3 5 3 4 3 1 5",
		"0 1 0 6 0 0 6 5 6 2 0 6 6 0 2 6 4 1 2 1 3 4 1 3 2 1 4 1 5 4 5 5 4 3 5 5 3 2 4 3 3 2",
		"0 1 2 5 6 6 6 3 1 5 6 6 2 3 3 5 4 3 5 2 3 3 1 0 4 2 1 1 1 4 6 0 0 0 0 5 5 4 4 4 2 2",
		"0 1 4 1 6 6 6 6 4 3 5 6 3 2 3 4 1 1 1 2 4 1 0 6 4 5 3 0 0 4 0 3 3 5 0 5 2 2 5 5 2 2",
		"0 1 6 1 2 2 2 5 5 3 3 6 5 0 2 0 2 3 1 2 4 1 4 1 6 4 0 6 1 6 6 3 4 3 0 0 4 4 3 5 5 5",
		"0 2 2 4 3 3 6 3 0 0 2 3 1 2 2 2 4 5 5 1 0 5 5 1 4 1 4 6 3 0 6 3 0 4 4 5 6 6 1 5 6 1",
		"0 2 3 2 3 4 1 5 2 0 1 6 5 2 3 4 2 2 6 3 5 0 0 6 6 4 0 3 4 0 1 5 5 4 5 1 6 4 1 1 3 6",
		"0 2 3 4 1 3 5 4 5 5 2 6 2 0 5 3 2 6 5 3 0 0 3 6 3 2 4 5 1 1 6 0 0 6 6 1 1 1 4 2 4 4",
		"0 2 5 1 3 6 4 3 0 1 2 0 4 6 4 4 3 6 6 5 6 2 4 6 3 1 0 2 3 5 1 4 5 3 2 1 2 1 0 5 5 0",
		"0 3 4 6 3 3 3 4 6 6 6 2 4 1 5 6 5 5 4 5 5 0 3 5 6 4 3 4 0 2 0 0 0 2 1 1 2 1 1 2 1 2",
		"0 3 6 0 3 3 3 4 4 6 4 5 2 4 2 3 3 5 4 2 0 1 2 2 5 5 5 0 6 2 6 0 5 1 4 0 1 1 1 6 1 6",
		"0 3 6 1 4 2 6 6 0 1 4 6 2 3 3 6 1 0 3 0 2 1 6 4 4 5 5 1 1 3 0 2 2 5 5 3 4 5 0 5 4 2",
		"0 4 0 3 1 3 4 1 2 3 3 0 2 1 6 1 3 5 6 6 1 0 2 2 3 1 4 0 0 5 6 5 6 4 5 4 2 4 2 6 5 5",
		"0 4 2 3 0 3 1 0 4 5 0 5 6 4 3 0 3 5 3 1 0 3 6 4 6 4 2 6 6 6 4 1 5 5 2 2 2 1 1 2 1 5",
		"0 4 6 2 1 0 5 4 6 5 3 3 0 5 4 0 6 3 1 0 5 1 2 5 1 6 1 2 2 0 1 4 3 5 4 6 6 2 2 4 3 3",
		"0 4 6 5 4 4 4 4 0 6 3 0 0 5 4 1 6 3 1 1 1 1 0 6 6 2 0 6 1 2 2 2 2 5 5 2 5 3 5 3 3 3",
		"0 5 2 3 5 0 0 1 4 1 2 5 1 4 1 3 3 3 5 3 5 5 6 1 6 0 0 1 0 2 4 6 3 4 6 6 2 6 4 2 2 4",
		"0 5 3 5 0 0 1 2 0 1 1 6 0 1 4 0 3 2 6 3 3 1 4 4 1 3 6 2 5 6 2 5 6 3 5 4 2 4 4 2 6 5",
		"0 5 4 5 5 3 3 0 1 5 5 5 0 2 4 1 4 3 1 0 3 3 3 2 2 4 6 6 4 0 6 0 2 6 6 4 6 1 1 1 2 2",
		"0 6 1 0 5 3 1 6 2 0 0 5 4 4 1 1 6 2 5 0 1 2 1 5 3 6 3 6 6 4 3 5 2 3 4 0 3 2 5 2 4 4",
		"0 6 3 5 4 4 3 2 5 2 6 6 2 3 6 0 1 5 4 5 0 0 5 6 4 3 4 0 5 1 1 4 0 1 1 1 2 2 2 6 3 3",
		"0 6 5 5 0 0 0 3 6 6 2 0 1 3 1 5 5 1 1 5 3 1 2 3 2 5 1 4 0 2 3 2 4 6 4 4 4 4 6 2 3 6",
		"0 6 6 3 2 3 2 6 6 0 5 2 3 0 2 3 1 4 2 6 0 6 1 5 5 1 1 1 2 5 1 5 5 0 3 4 0 3 4 4 4 4",
		"0 6 6 4 5 6 6 3 6 6 1 5 0 4 1 0 0 5 2 5 4 2 3 0 4 0 2 3 1 4 5 3 4 3 3 1 2 5 1 2 2 1",
		"1 0 2 3 2 2 4 6 2 5 6 1 1 1 6 4 0 5 3 3 4 4 6 2 1 0 0 0 1 2 5 5 5 6 0 4 3 3 3 4 6 5",
		"1 0 4 1 4 4 5 3 3 3 1 6 3 2 2 1 4 4 3 5 1 1 0 5 2 0 0 6 4 2 6 2 2 6 3 0 0 5 5 5 6 6",
		"1 0 5 1 6 1 1 3 4 5 2 0 1 1 4 5 2 3 5 4 4 4 0 3 0 4 0 0 5 5 2 2 3 2 3 2 3 6 6 6 6 6",
		"1 0 6 5 3 3 2 1 0 2 6 4 4 0 5 1 2 3 1 5 1 1 0 5 6 2 5 2 5 2 3 4 0 6 3 4 6 4 4 6 3 0",
		"1 0 6 6 3 1 0 4 4 5 1 5 4 6 1 1 0 4 2 0 3 6 0 1 2 3 3 2 6 2 0 3 4 2 6 4 5 3 2 5 5 5",
		"1 1 1 0 5 5 0 0 1 2 1 6 5 2 5 5 5 4 2 6 0 0 0 3 3 4 3 2 3 3 6 6 2 4 4 1 6 2 3 4 4 6",
		"1 1 3 2 6 1 2 3 3 5 2 3 2 6 4 2 3 3 0 2 4 0 1 5 6 1 0 5 4 4 1 0 5 4 0 5 5 6 4 6 6 0",
		"1 1 6 5 4 2 2 2 4 1 4 1 1 5 6 5 1 6 5 6 3 6 3 3 6 4 0 3 5 5 3 3 4 0 0 4 0 0 2 2 2 0",
		"1 2 2 0 5 5 4 0 0 1 3 3 0 1 3 1 1 5 2 2 0 0 4 1 4 4 3 5 5 2 3 4 2 6 4 6 5 3 6 6 6 6",
		"1 2 5 3 4 2 0 0 2 5 0 4 3 2 0 6 6 6 2 3 4 4 5 5 5 1 1 4 3 3 2 0 3 0 5 4 6 6 6 1 1 1",
		"1 2 5 4 2 5 3 3 0 6 2 4 2 1 6 6 0 2 6 1 0 6 4 3 3 4 6 0 5 4 3 5 3 5 4 0 0 2 1 5 1 1",
		"1 3 4 4 1 3 0 6 6 4 3 1 2 5 6 0 6 1 0 0 2 4 4 2 3 2 3 3 0 2 1 5 2 4 1 0 5 6 5 5 5 6",
		"1 4 1 5 6 0 1 0 0 3 3 6 6 1 0 0 0 3 4 4 6 5 3 6 2 2 2 1 6 4 3 1 5 4 3 5 4 2 5 5 2 2",
		"1 4 2 6 5 0 3 0 5 6 4 4 1 5 3 3 6 3 3 3 0 4 0 2 5 1 4 6 2 0 2 4 5 2 2 1 0 1 1 6 5 6",
		"1 4 5 3 2 6 6 5 6 6 4 4 6 6 3 3 3 4 0 1 0 3 5 4 5 2 5 5 1 2 2 1 3 2 0 0 1 0 4 2 0 1",
		"1 4 6 6 5 1 1 2 0 4 1 5 4 0 5 6 1 0 3 0 0 0 4 6 6 1 6 5 4 4 3 5 5 3 2 2 2 3 3 2 3 2",
		"1 5 0 2 5 5 5 6 3 6 1 4 4 1 6 5 6 3 4 4 5 4 2 0 0 4 0 6 2 3 1 1 6 3 1 2 3 3 2 0 0 2",
		"1 5 6 2 2 2 5 3 4 5 6 3 4 0 1 1 3 6 5 5 3 2 3 0 6 1 6 2 1 6 0 5 0 4 0 4 1 0 2 4 4 3",
		"1 6 4 0 3 2 5 1 5 4 2 2 0 2 4 2 0 5 1 1 5 3 1 3 3 1 2 4 3 6 5 0 3 5 0 6 6 4 0 6 6 4",
		"1 6 4 4 5 5 2 0 2 3 3 4 3 6 5 4 3 0 1 3 0 5 1 1 1 0 5 3 0 6 5 2 4 4 0 1 6 2 6 2 2 6",
		"1 6 6 2 6 0 6 3 2 0 1 1 0 6 4 3 4 3 1 2 4 0 2 0 3 2 3 6 0 2 1 4 3 1 5 4 4 5 5 5 5 5",
		"2 0 0 5 6 5 2 0 4 3 2 6 1 6 5 5 0 0 4 1 0 2 6 4 6 2 6 5 1 4 4 5 1 4 1 3 3 1 3 3 2 3",
		"2 1 6 3 4 2 0 4 5 2 5 0 3 4 3 0 3 3 5 5 1 2 0 5 4 3 6 6 5 0 1 0 4 6 2 1 4 1 6 2 1 6",
		"2 2 0 0 5 6 2 1 1 0 0 4 5 5 3 6 4 4 6 1 4 2 2 3 5 4 3 2 5 6 1 1 5 0 0 4 6 3 1 6 3 3",
		"2 2 0 2 2 2 2 0 0 5 3 4 6 4 3 6 6 1 0 0 1 4 4 6 6 4 5 5 3 3 4 6 1 1 1 3 1 5 5 3 5 0",
		"2 2 4 5 2 1 5 5 2 4 5 6 1 1 0 0 3 3 0 1 3 0 0 1 5 0 6 2 6 4 5 3 1 6 3 6 6 3 2 4 4 4",
		"2 2 5 5 5 3 6 0 0 4 2 3 3 0 5 6 1 3 4 4 1 4 4 2 4 3 2 5 0 5 0 0 3 6 6 1 1 2 6 1 6 1",
		"2 3 0 5 4 6 2 1 5 1 2 2 4 6 0 6 0 2 5 0 5 5 1 2 4 4 4 3 0 4 1 3 6 5 1 0 3 3 3 6 6 1",
		"2 3 2 1 3 1 5 5 6 2 5 0 0 5 4 3 3 6 6 5 4 1 2 2 0 0 4 3 0 5 3 0 1 2 6 1 1 6 6 4 4 4",
		"2 3 3 4 1 4 6 0 4 3 1 5 6 0 1 4 0 2 0 1 3 5 4 6 3 1 5 1 0 4 6 5 5 0 3 2 2 6 6 2 5 2",
		"2 3 4 4 2 5 3 0 5 4 0 0 3 4 3 3 4 4 5 6 6 6 0 6 6 6 0 1 0 5 5 2 5 1 1 2 1 1 2 3 2 1",
		"2 4 4 3 6 6 6 4 1 6 0 1 3 5 1 0 4 4 5 2 5 0 2 5 0 0 6 2 2 1 6 5 0 3 4 5 3 1 1 2 3 3",
		"2 4 4 5 1 6 5 4 1 1 4 0 5 0 1 2 2 1 3 4 2 6 0 3 2 5 4 5 5 2 0 0 3 1 6 6 6 6 0 3 3 3",
		"2 5 5 1 3 2 0 0 0 2 2 6 1 1 4 3 0 4 2 1 0 6 6 1 6 3 1 2 5 3 3 0 4 6 6 3 4 5 5 4 5 4",
		"2 6 1 0 0 1 5 3 6 6 1 6 6 3 4 4 6 3 2 4 3 0 4 1 0 3 3 5 5 2 5 2 1 1 0 5 0 4 5 4 2 2",
		"2 6 4 1 3 1 6 3 0 6 1 1 4 4 4 1 4 5 4 6 3 3 0 0 5 6 1 2 0 3 6 5 2 3 0 0 5 2 2 5 2 5",
		"2 6 5 6 5 0 1 6 0 4 4 5 5 2 5 0 2 0 4 2 3 2 0 3 0 4 1 5 6 2 6 1 3 4 1 6 1 3 3 1 3 4",
		"3 1 5 2 4 0 3 6 2 3 6 0 4 1 0 6 4 0 1 5 3 4 2 1 2 1 4 1 0 6 6 2 5 3 4 5 6 2 3 0 5 5",
		"3 1 5 6 6 2 4 6 0 5 3 3 1 0 0 1 2 4 4 4 4 2 0 3 5 0 5 0 3 3 4 1 2 2 6 1 2 6 1 6 5 5",
		"3 1 6 5 2 0 6 6 5 2 3 5 6 2 0 5 5 4 3 6 0 4 0 6 1 5 4 1 2 0 1 3 3 1 3 1 4 2 0 2 4 4",
		"3 1 6 6 0 5 6 6 6 0 6 1 2 4 4 5 1 3 4 5 2 4 0 3 5 4 3 1 1 4 3 3 1 0 0 0 2 2 5 2 5 2",
		"3 2 1 5 1 3 3 3 2 2 3 1 1 6 2 4 6 0 0 5 3 4 0 1 4 6 6 4 4 0 4 0 5 6 5 2 1 2 6 5 0 5",
		"3 2 2 0 1 6 1 0 5 1 3 3 4 1 0 4 1 1 4 5 2 3 5 5 2 3 6 2 0 6 3 0 6 2 6 6 4 5 5 0 4 4",
		"3 2 4 5 6 4 5 4 5 1 6 6 3 4 3 2 6 3 1 0 5 0 4 3 6 2 6 0 2 5 2 5 2 4 1 1 0 0 3 0 1 1",
		"3 2 6 5 0 4 6 1 6 5 4 5 2 6 5 2 4 2 1 5 4 6 5 0 6 1 1 3 1 0 1 0 0 0 3 3 3 4 3 2 2 4",
		"3 3 0 1 0 6 5 5 6 4 6 4 2 3 4 6 2 4 5 2 5 4 3 4 2 5 5 2 2 6 0 1 3 1 1 0 6 3 0 1 0 1",
		"3 3 0 2 0 4 5 0 6 0 2 1 4 4 5 6 3 3 0 3 0 5 1 1 6 2 5 6 1 3 6 2 4 6 2 1 2 5 4 5 4 1",
		"3 3 1 6 4 3 1 5 3 2 3 1 3 0 4 1 5 6 1 2 2 2 6 1 0 6 0 4 0 2 4 4 5 2 6 6 5 5 4 0 5 0",
		"3 3 3 5 5 1 6 1 0 6 0 2 2 0 5 1 5 5 5 3 4 0 0 4 1 2 2 0 2 3 4 1 4 3 4 6 6 2 1 4 6 6",
		"3 4 4 5 2 0 6 2 4 6 1 3 6 4 2 3 2 2 1 6 1 6 6 5 5 0 3 4 4 3 0 5 3 2 5 0 0 1 1 5 1 0",
		"3 4 5 0 3 0 0 6 5 6 1 2 0 4 1 3 3 4 3 1 5 2 4 5 3 6 5 0 0 4 2 1 6 4 1 6 1 6 5 2 2 2",
		"3 6 6 1 3 6 4 0 6 2 1 2 0 2 0 4 5 1 2 6 1 6 5 4 5 5 5 2 0 4 4 2 3 5 4 3 1 0 3 0 1 3",
		"3 6 6 4 1 4 4 2 1 2 2 2 5 5 3 6 1 1 1 1 2 6 2 0 0 3 5 6 3 0 5 4 5 4 0 0 4 5 0 3 6 3",
		"4 0 1 3 0 0 1 2 2 5 0 6 2 3 4 0 6 1 4 4 5 0 5 4 4 3 1 6 6 6 3 2 1 2 1 5 6 5 3 2 3 5",
		"4 0 3 1 1 1 6 0 0 5 6 6 6 2 5 4 1 4 1 6 6 4 2 3 4 0 3 5 2 0 5 5 5 3 4 2 1 2 3 0 2 3",
		"4 0 3 6 0 5 1 6 5 3 3 3 3 5 4 2 6 1 2 0 2 3 5 5 6 0 5 6 0 1 0 2 1 2 1 4 1 6 2 4 4 4",
		"4 1 5 3 6 1 1 1 3 6 2 0 6 4 0 1 2 0 1 4 2 6 3 6 4 2 4 3 5 3 5 4 0 5 6 5 2 2 5 3 0 0",
		"4 1 6 3 5 5 6 3 3 5 1 4 5 1 2 6 2 5 4 4 4 5 0 0 1 2 4 3 2 1 6 0 3 1 3 0 0 0 2 6 6 2",
		"4 2 0 5 1 3 0 1 6 3 5 5 4 0 3 2 1 3 6 6 2 4 4 1 0 4 3 1 1 3 5 2 4 5 6 6 6 5 2 2 0 0",
		"4 4 0 5 5 4 5 2 2 6 3 1 5 5 0 6 6 0 1 6 2 1 6 4 5 3 3 2 4 2 1 3 6 1 3 2 4 3 1 0 0 0",
		"4 4 1 3 3 3 5 5 2 1 3 0 0 2 1 3 6 1 2 2 4 5 4 5 5 5 0 4 6 6 1 4 3 2 6 2 0 1 6 0 6 0",
		"4 4 2 5 2 4 0 2 3 6 5 1 3 2 6 0 6 1 3 3 3 1 1 3 2 0 4 5 5 0 2 5 6 1 0 0 4 6 1 4 6 5",
		"4 4 6 1 0 5 0 0 0 2 4 2 5 2 2 6 0 0 3 3 1 4 4 1 3 3 2 3 2 4 3 6 1 1 5 1 6 6 5 5 5 6",
		"4 5 3 2 6 5 0 6 3 2 5 2 0 0 1 6 6 6 0 6 3 1 1 4 2 5 0 0 4 2 4 3 4 1 5 5 1 4 2 3 3 1",
		"4 5 4 0 0 4 4 6 2 3 1 2 6 3 2 1 6 4 5 0 1 3 2 5 3 0 0 3 2 1 0 2 1 5 3 1 4 5 5 6 6 6",
		"4 5 6 4 2 1 6 0 0 2 3 6 5 3 1 3 4 2 4 1 0 5 2 4 6 2 1 6 5 4 1 6 5 3 5 1 0 0 2 0 3 3",
		"4 6 4 0 4 2 0 2 5 2 6 6 2 5 0 2 6 5 2 5 5 3 3 0 0 3 3 4 4 3 0 3 1 5 4 6 1 6 1 1 1 1",
		"4 6 4 1 0 0 6 3 5 3 2 3 3 5 2 2 3 6 4 4 6 2 4 6 6 1 0 0 1 5 5 0 5 5 0 4 1 2 2 1 1 3",
		"4 6 6 1 2 5 4 5 1 2 5 6 1 0 0 4 4 0 5 4 3 1 4 3 0 2 3 2 2 6 6 2 6 5 3 1 0 3 0 3 5 1",
		"5 0 3 1 3 3 6 6 6 6 2 1 6 5 3 2 5 3 1 5 3 4 6 5 2 2 0 0 0 1 1 4 4 2 4 5 2 0 4 4 1 0",
		"5 0 6 2 1 4 1 1 5 6 0 0 3 5 5 0 6 3 4 4 6 0 0 2 2 4 6 4 4 5 1 1 5 1 3 2 2 3 2 6 3 3",
		"5 1 1 6 5 1 2 1 1 4 1 6 4 0 0 5 0 2 6 4 2 6 0 5 4 5 5 2 6 0 0 3 2 3 2 6 3 4 3 4 3 3",
		"5 1 6 5 6 3 5 4 0 0 4 5 5 0 2 5 0 4 2 0 3 1 3 2 0 4 1 1 6 6 3 3 3 6 1 2 2 6 2 4 4 1",
		"5 2 2 4 2 3 1 3 6 6 0 1 1 6 1 3 4 5 5 2 1 1 3 2 4 4 0 3 3 0 6 2 6 0 4 4 6 0 0 5 5 5",
		"5 3 1 3 3 6 5 5 3 2 6 3 1 1 5 6 2 2 0 3 2 5 4 2 0 4 1 1 2 0 1 4 6 4 4 4 5 0 0 0 6 6",
		"5 3 3 5 3 1 6 6 5 0 2 6 4 1 6 0 0 6 0 5 1 5 5 6 3 3 2 3 0 2 1 2 4 0 2 4 2 1 1 4 4 4",
		"5 3 4 0 6 2 3 6 2 5 1 3 2 4 0 1 2 2 3 1 0 1 6 6 6 6 3 2 4 3 1 4 5 5 1 0 0 4 5 5 4 0",
		"5 3 5 0 0 0 1 0 5 1 2 5 3 4 6 3 3 3 0 2 4 4 0 4 5 5 1 1 1 6 6 1 4 2 3 6 2 2 4 2 6 6",
		"5 3 6 3 6 6 4 1 1 3 1 6 2 1 3 3 6 1 0 6 5 4 1 2 0 4 5 3 4 2 2 5 4 2 0 5 4 0 0 5 0 2",
		"5 4 0 0 2 1 4 6 3 3 0 6 4 0 6 5 4 4 1 4 0 6 6 6 0 1 5 3 5 5 1 3 3 3 5 1 2 2 1 2 2 2",
		"5 4 1 1 5 6 2 0 3 1 4 0 2 6 4 0 1 3 4 5 6 4 1 3 2 3 4 2 0 2 3 0 2 3 5 1 5 5 0 6 6 6",
		"5 4 3 5 0 6 6 3 4 6 4 0 2 0 0 1 4 5 6 2 0 6 1 6 2 2 3 0 5 4 2 1 1 1 5 3 4 1 3 2 3 5",
		"5 4 4 0 1 2 0 4 1 1 4 4 0 0 1 0 3 1 5 2 2 3 4 1 3 5 6 3 3 2 2 3 5 2 0 6 6 5 6 5 6 6",
		"5 4 4 0 3 0 4 3 3 2 4 4 6 2 5 5 0 1 1 4 5 1 0 3 3 6 0 1 3 5 6 5 6 1 2 6 6 0 1 2 2 2",
		"5 4 4 2 3 6 1 3 4 1 0 1 2 0 3 3 0 2 3 5 6 2 0 2 1 6 2 6 0 4 4 5 4 0 5 1 1 5 5 3 6 6",
		"5 5 4 6 1 1 0 5 6 3 1 3 6 1 4 4 2 1 3 0 6 4 4 6 0 0 3 4 0 1 2 2 5 3 0 5 6 2 5 2 2 3",
		"5 6 4 4 1 2 6 0 2 3 4 4 4 0 4 6 5 5 6 2 2 2 0 6 0 6 1 0 1 5 3 1 3 3 2 0 5 5 1 3 1 3",
		"5 6 6 3 1 6 3 5 5 2 5 2 6 3 3 6 0 2 2 2 5 6 4 1 2 3 1 0 4 1 1 5 4 1 0 4 0 4 4 3 0 0",
		"5 6 6 6 2 0 0 3 0 0 2 6 4 3 0 0 2 2 1 1 6 1 4 1 4 5 6 3 1 2 2 5 1 5 5 4 3 5 3 4 3 4",
		"6 0 1 4 0 1 4 2 5 1 1 2 3 0 0 5 0 4 0 3 2 5 4 6 3 2 2 2 5 1 6 6 3 5 4 1 3 3 6 6 4 5",
		"6 0 3 0 3 1 0 2 5 2 5 5 5 6 0 4 3 3 1 4 6 2 3 5 4 4 0 3 2 2 4 1 2 5 1 4 6 1 1 6 6 0",
		"6 1 2 3 5 0 2 1 1 5 5 5 1 0 3 4 1 1 2 2 4 3 3 3 0 2 0 6 4 5 3 5 6 6 2 0 6 0 6 4 4 4",
		"6 1 4 4 5 5 0 0 4 2 2 3 5 5 1 2 3 0 2 5 6 3 0 4 2 5 4 2 1 6 1 1 4 1 6 6 0 3 6 3 3 0",
		"6 1 5 1 0 4 4 2 5 5 2 0 3 1 0 4 2 2 5 5 1 1 2 4 5 3 1 6 3 4 2 0 3 6 4 0 6 0 3 6 6 3",
		"6 2 1 1 3 1 1 5 0 2 6 1 3 4 5 1 6 2 3 4 5 5 2 6 6 4 6 0 2 0 0 2 0 5 4 0 4 3 4 3 3 5",
		"6 2 2 4 0 6 5 6 6 1 1 4 0 2 4 1 1 0 2 4 3 6 4 1 1 2 0 2 0 3 0 6 5 5 4 5 5 5 3 3 3 3",
		"6 2 3 0 1 3 0 6 4 1 4 5 5 4 6 0 2 1 4 3 2 2 3 6 3 3 1 5 5 4 6 6 0 4 1 5 5 0 1 2 2 0",
		"6 2 3 4 0 1 2 5 1 3 6 1 6 3 2 4 4 3 2 5 3 6 1 6 4 2 2 3 6 4 0 0 1 5 5 4 5 5 0 0 1 0",
		"6 2 5 1 5 4 3 3 6 0 0 4 0 5 3 6 0 2 4 3 6 0 1 5 3 4 2 1 2 5 5 0 6 6 3 1 2 4 4 2 1 1",
		"6 3 0 4 6 0 2 1 5 6 3 6 6 6 3 5 2 1 1 0 5 1 5 5 1 5 2 0 0 2 4 4 3 3 2 3 2 4 1 4 4 0",
		"6 3 1 4 3 6 2 2 5 3 6 0 4 4 3 1 0 6 0 5 4 3 1 1 2 0 0 6 2 1 4 2 6 2 3 0 4 1 5 5 5 5",
		"6 3 2 1 3 2 5 0 5 0 3 1 4 6 0 0 0 0 5 4 3 3 6 5 1 1 3 2 2 6 2 1 5 1 5 2 6 4 6 4 4 4",
		"6 3 5 6 3 5 0 2 1 6 1 3 3 6 1 1 5 5 5 3 3 1 4 0 6 4 2 1 4 4 5 2 4 4 6 0 2 2 2 0 0 0",
		"6 4 1 2 0 3 2 6 0 1 5 3 4 0 4 3 6 5 1 1 0 5 3 1 2 6 4 0 0 4 5 3 5 5 3 6 6 1 4 2 2 2",
		"6 4 2 2 3 5 6 1 4 0 1 3 4 6 5 3 0 5 4 3 3 6 6 4 6 0 3 1 0 0 5 0 2 2 1 5 1 2 1 2 5 4",
		"6 4 4 5 6 0 6 2 1 5 3 3 0 6 3 3 4 4 1 5 2 1 5 0 2 6 5 0 3 5 0 4 6 2 4 3 1 0 1 2 1 2",
		"6 5 1 3 2 4 2 5 6 6 4 4 5 0 4 0 0 5 6 4 0 4 0 1 3 3 1 1 6 3 6 5 1 5 1 0 2 2 2 3 3 2",
		"6 6 1 6 0 2 0 4 3 1 2 0 0 3 6 4 3 1 5 6 4 6 2 3 5 4 2 5 5 5 3 0 1 1 5 3 0 4 4 2 2 1",
		"6 6 2 2 4 6 5 0 2 2 0 5 4 4 6 4 6 1 4 0 6 3 2 1 1 3 5 1 0 5 0 5 3 1 3 5 4 0 2 3 1 3",
		"6 6 4 3 1 5 3 5 0 2 4 1 1 6 1 6 0 2 5 2 6 5 2 0 3 0 0 2 3 1 6 0 1 5 2 3 5 3 4 4 4 4",
	}), shuffle([]string{
		"0 0 1 1 0 1 5 5 2 3 6 3 2 0 4 5 2 4 5 2 3 0 6 5 0 2 4 2 1 6 3 4 3 1 1 5 6 4 6 4 6 3",
		"0 0 3 0 4 1 1 6 5 3 4 3 5 6 4 2 6 4 0 3 6 5 3 1 1 4 1 5 2 3 2 5 2 6 0 0 1 2 6 2 4 5",
		"0 1 6 2 4 5 4 0 0 1 3 0 5 3 2 2 5 0 5 6 4 6 0 3 2 5 2 1 2 3 1 4 3 1 1 3 4 4 5 6 6 6",
		"0 2 0 1 3 0 3 2 6 2 3 3 2 2 1 5 4 2 5 4 3 1 5 6 1 3 5 5 5 6 6 1 4 0 4 4 4 6 6 1 0 0",
		"0 2 1 2 3 1 3 0 4 4 1 2 6 1 3 0 1 5 2 2 5 6 1 4 0 3 5 2 0 3 6 0 6 4 5 6 4 4 6 5 5 3",
		"0 2 1 4 2 0 2 1 4 3 1 3 4 3 3 1 3 0 6 3 1 4 0 1 6 0 0 4 4 6 5 2 6 5 2 5 5 5 5 2 6 6",
		"0 2 4 0 6 1 6 2 1 0 0 1 0 3 1 6 0 5 3 1 3 2 2 3 6 3 6 5 4 4 1 3 4 4 2 5 5 5 4 2 6 5",
		"0 2 5 6 0 0 3 6 0 4 0 2 6 5 6 3 4 0 2 5 3 3 1 5 6 2 2 6 4 3 2 1 5 1 4 4 1 5 1 1 3 4",
		"0 3 2 3 3 0 2 0 3 3 4 0 5 6 5 3 6 6 5 6 4 4 2 5 6 4 0 0 5 2 4 6 2 2 4 5 1 1 1 1 1 1",
		"0 3 3 2 6 3 1 1 6 6 3 5 1 3 2 5 4 2 0 4 0 4 4 6 3 0 4 1 5 1 2 1 0 2 2 4 0 5 5 6 5 6",
		"0 3 4 2 2 2 3 5 0 4 1 4 3 1 5 4 4 3 1 0 2 6 2 2 6 6 6 1 5 6 5 3 1 5 5 3 0 1 6 0 4 0",
		"0 3 4 6 2 4 0 1 6 6 6 5 5 5 0 5 6 0 0 5 2 3 3 1 0 2 5 6 3 1 4 2 4 2 1 4 2 1 1 4 3 3",
		"0 4 1 2 1 6 0 2 5 3 2 6 6 2 5 6 3 1 6 4 5 1 1 5 2 1 2 5 3 3 5 3 6 0 0 4 4 4 0 3 0 4",
		"0 4 6 2 5 4 4 2 6 6 2 1 4 4 0 6 5 2 1 2 2 5 6 4 3 0 1 5 0 5 6 0 0 1 5 3 1 3 1 3 3 3",
		"0 5 1 6 5 1 3 2 6 3 0 3 3 1 3 3 6 5 5 4 4 0 5 5 1 6 2 6 6 4 2 2 0 1 4 2 2 4 4 0 1 0",
		"0 6 4 6 6 5 3 6 4 1 2 6 5 3 5 1 1 2 3 4 5 1 2 2 4 6 4 1 4 1 3 5 2 2 0 0 0 5 3 3 0 0",
		"0 6 5 0 4 3 6 3 4 0 0 2 4 6 1 1 2 4 2 6 6 6 5 5 1 5 5 3 3 1 5 4 4 3 0 2 3 0 1 2 1 2",
		"0 6 6 2 1 6 5 1 3 0 6 2 3 6 1 0 2 6 5 4 1 4 4 5 5 1 0 4 5 3 4 5 2 3 4 0 0 1 2 3 3 2",
		"0 6 6 4 4 6 4 5 3 6 5 5 6 1 5 4 0 3 2 2 3 5 5 6 4 0 0 4 0 1 2 1 1 1 1 0 2 3 2 3 3 2",
		"1 0 0 1 2 3 6 1 2 5 3 1 1 3 2 2 2 2 3 3 6 0 6 6 4 0 3 5 0 6 5 6 4 4 4 0 1 5 5 5 4 4",
		"1 0 4 6 4 3 2 5 0 5 1 1 6 0 5 4 2 3 6 3 5 4 5 6 3 4 4 0 2 5 6 2 0 2 2 6 1 1 3 3 1 0",
		"1 1 1 4 4 3 5 1 1 6 0 2 2 3 0 5 5 6 4 3 0 4 2 6 3 3 4 0 3 2 2 1 4 5 2 0 0 5 6 5 6 6",
		"1 1 4 3 5 0 1 0 1 6 4 5 5 6 1 3 2 4 2 3 6 4 5 2 0 1 3 2 0 4 3 0 0 5 2 3 4 6 5 6 2 6",
		"1 1 5 2 0 1 0 0 6 1 3 4 1 0 1 6 6 4 2 4 5 6 2 6 3 3 3 5 4 4 4 5 2 6 5 2 5 0 0 2 3 3",
		"1 1 5 5 0 1 5 5 2 6 0 4 5 4 6 3 2 1 1 6 0 5 1 0 3 6 2 3 6 2 4 0 3 0 4 4 6 3 4 2 2 3",
		"1 2 2 0 2 5 4 4 1 6 6 1 4 4 3 5 6 0 2 2 6 4 2 6 4 6 3 1 5 5 5 0 0 3 1 5 1 3 3 0 0 3",
		"1 2 4 4 4 6 0 5 0 4 1 6 5 3 4 0 1 6 6 2 0 4 3 2 0 0 6 6 3 3 2 1 2 3 2 1 1 3 5 5 5 5",
		"1 2 5 1 6 4 2 6 1 2 4 5 5 1 5 4 0 4 1 6 4 3 5 6 2 0 1 0 3 4 2 3 3 5 6 3 2 0 0 6 0 3",
		"1 2 6 6 0 0 1 2 2 4 5 0 6 1 0 4 6 3 5 0 0 1 2 2 4 1 2 4 1 3 3 5 5 5 4 3 5 4 3 6 6 3",
		"1 3 1 2 6 1 3 1 1 0 2 4 1 3 6 0 3 6 0 4 4 2 6 4 2 6 0 2 5 6 3 3 4 4 2 5 0 0 5 5 5 5",
		"1 3 3 5 0 5 2 3 5 1 1 5 2 4 4 5 6 3 5 2 0 6 6 4 1 4 6 3 3 0 4 0 0 0 2 6 1 1 2 4 2 6",
		"1 3 3 6 4 6 4 1 3 4 6 5 3 5 6 4 4 3 1 6 0 0 5 1 0 0 6 3 0 1 4 1 2 0 5 5 2 2 2 2 5 2",
		"1 3 5 6 3 0 2 0 0 1 2 0 1 3 4 0 6 0 6 4 1 1 5 1 4 4 6 2 2 2 3 6 2 6 4 4 3 3 5 5 5 5",
		"1 3 6 5 0 1 2 2 1 4 6 4 4 4 1 6 5 5 3 4 1 1 6 4 2 0 0 3 0 5 2 2 5 0 6 5 2 3 6 0 3 3",
		"1 4 0 0 1 3 2 4 4 1 3 2 3 1 6 2 0 4 6 1 1 3 0 6 6 4 0 6 5 0 6 3 2 4 3 5 2 2 5 5 5 5",
		"1 4 2 2 1 3 1 2 5 1 3 3 0 6 2 4 4 0 1 2 3 5 5 5 6 3 3 2 5 4 1 0 0 6 0 6 4 4 5 6 0 6",
		"1 4 6 6 0 6 2 1 5 3 1 2 0 1 6 2 1 6 4 0 6 5 3 4 2 3 2 2 0 0 3 3 1 3 4 0 4 4 5 5 5 5",
		"1 5 0 3 1 1 1 3 1 1 4 6 6 6 0 6 3 0 0 3 2 0 5 2 2 6 5 4 5 5 0 3 6 2 5 3 2 4 2 4 4 4",
		"1 5 0 3 6 6 5 5 5 0 6 6 4 5 6 0 5 3 2 2 0 0 4 4 0 2 1 1 3 2 2 4 1 2 6 1 4 1 3 4 3 3",
		"1 6 0 6 4 0 3 2 4 4 3 1 6 4 1 5 0 0 2 0 0 6 4 3 3 2 2 3 2 1 4 5 5 3 5 5 6 2 5 1 1 6",
		"1 6 4 5 3 5 4 2 6 0 2 0 4 2 0 3 6 5 2 4 5 6 5 5 6 6 1 4 3 0 4 1 0 1 3 1 2 3 3 0 2 1",
		"2 0 0 4 0 0 3 4 0 0 2 2 6 3 2 6 4 2 1 2 6 4 3 4 6 4 1 6 6 3 3 5 5 1 1 5 3 1 1 5 5 5",
		"2 0 4 3 1 1 4 6 6 3 5 6 6 6 6 3 3 1 2 5 5 2 3 3 0 1 4 4 4 5 2 4 1 2 0 0 0 2 5 0 1 5",
		"2 1 0 6 0 2 1 1 3 0 0 3 1 1 2 1 5 2 3 6 0 6 0 2 6 4 4 6 4 6 4 4 2 3 5 5 3 3 5 4 5 5",
		"2 1 1 6 1 0 2 0 0 2 4 2 4 5 3 0 1 2 4 4 4 6 2 1 0 3 6 4 5 3 1 3 0 5 3 3 5 6 6 5 5 6",
		"2 2 1 3 4 1 0 5 6 6 2 0 2 2 1 2 6 4 5 0 6 4 5 6 3 4 3 5 6 3 0 5 4 0 3 4 3 1 5 1 0 1",
		"2 2 4 3 1 0 4 5 1 3 4 5 5 1 0 5 5 0 6 5 2 2 2 6 6 4 0 0 2 3 6 6 6 1 3 4 4 3 3 1 0 1",
		"2 2 5 4 5 5 0 1 6 2 4 0 0 5 5 6 2 0 5 2 4 6 4 6 6 0 1 4 6 0 4 2 3 1 1 3 1 3 1 3 3 3",
		"2 3 0 0 5 0 0 2 1 2 2 5 1 1 3 6 6 1 4 4 4 6 6 6 2 0 3 3 1 0 1 3 5 5 5 6 2 5 4 3 4 4",
		"2 3 0 1 1 6 4 5 4 6 0 3 1 4 2 1 0 2 5 0 1 3 0 0 3 3 4 3 2 2 4 4 5 5 5 1 2 5 6 6 6 6",
		"2 3 1 4 6 4 6 3 0 0 4 4 1 4 0 6 4 6 6 0 0 1 3 0 3 6 2 5 1 2 1 3 2 5 2 1 5 5 2 5 5 3",
		"2 3 2 3 4 6 5 4 5 3 1 2 3 4 3 0 2 3 5 5 2 1 4 6 1 0 5 2 0 4 5 1 6 0 6 4 0 1 6 6 0 1",
		"2 3 3 3 1 0 1 5 6 2 0 6 6 0 2 4 0 2 2 1 4 0 6 5 1 2 3 1 5 3 6 6 3 5 5 1 4 5 4 4 4 0",
		"2 4 1 3 0 3 5 1 6 0 2 5 2 5 4 6 0 3 3 4 5 2 2 4 2 4 4 1 3 3 0 0 6 1 1 1 6 5 0 5 6 6",
		"2 4 4 1 6 5 0 2 0 2 6 2 2 5 3 4 4 0 0 5 6 0 1 2 4 6 4 0 6 6 5 1 1 3 3 5 3 1 5 1 3 3",
		"2 4 5 1 5 3 0 0 6 3 3 5 2 0 5 4 1 6 0 1 4 6 0 2 4 3 0 4 1 3 6 4 2 5 3 1 6 1 2 2 6 5",
		"2 4 6 3 3 4 5 5 2 5 2 3 1 6 3 0 0 1 4 4 6 5 6 1 1 2 3 0 5 0 3 0 1 5 4 4 1 6 6 2 2 0",
		"2 5 5 0 1 2 0 5 5 6 2 0 0 4 5 4 6 0 2 2 6 6 6 4 3 0 2 6 3 1 3 3 4 3 4 1 4 5 1 3 1 1",
		"2 5 5 5 2 2 0 6 2 3 1 3 2 1 4 6 3 2 6 6 6 0 5 3 0 6 3 3 4 5 0 5 4 4 0 0 1 1 4 1 1 4",
		"2 6 2 1 3 3 1 4 6 6 0 6 5 1 6 2 3 4 0 0 3 5 3 6 1 2 0 3 4 4 1 0 1 5 5 4 0 2 4 5 5 2",
		"2 6 6 6 3 0 0 4 5 4 4 3 2 1 6 0 6 2 0 2 3 3 0 6 1 4 3 3 4 5 5 1 1 1 1 2 0 4 5 5 2 5",
		"3 0 1 6 6 1 0 5 4 3 5 0 4 4 1 2 6 2 6 3 0 4 2 2 1 0 4 6 0 4 6 5 3 1 5 2 1 2 5 5 3 3",
		"3 0 4 3 2 1 4 5 5 3 3 5 3 1 5 6 3 4 2 6 0 5 1 5 6 6 2 2 6 6 4 0 2 4 1 1 2 1 4 0 0 0",
		"3 0 6 2 1 1 0 3 5 5 5 6 2 0 1 3 3 4 4 0 1 5 6 6 6 3 1 2 5 1 0 5 2 2 0 3 2 6 4 4 4 4",
		"3 0 6 2 3 0 0 4 3 3 3 2 0 5 3 2 5 4 2 1 6 6 0 4 6 5 1 2 2 6 6 0 4 5 1 5 4 4 5 1 1 1",
		"3 1 2 5 0 3 0 4 2 5 6 5 6 0 1 4 3 1 3 3 1 1 3 6 5 2 5 2 2 1 0 0 6 0 4 6 5 4 6 4 2 4",
		"3 1 2 5 1 4 6 2 1 0 5 5 4 2 5 6 4 6 2 3 3 3 2 5 4 0 3 0 1 4 0 4 0 1 2 0 6 6 1 6 3 5",
		"3 1 4 6 0 5 1 2 1 2 5 0 0 5 2 1 5 0 1 1 3 4 0 5 2 2 4 3 0 5 6 3 3 2 4 6 6 3 4 6 6 4",
		"3 2 0 1 4 6 2 4 5 0 0 2 4 1 3 3 2 3 6 2 1 4 5 0 6 5 4 6 0 0 5 4 1 1 3 3 1 2 5 6 6 5",
		"3 2 0 2 1 5 1 3 5 5 5 1 3 6 6 3 2 4 4 6 5 6 6 4 2 3 3 4 0 0 6 2 2 1 4 0 5 4 0 1 0 1",
		"3 2 4 1 6 5 2 6 0 2 2 3 1 5 2 5 3 2 5 5 4 3 3 3 1 5 6 4 1 0 6 6 0 1 1 0 6 4 0 0 4 4",
		"3 2 6 2 6 5 1 2 4 4 3 4 5 6 4 6 2 6 4 0 6 0 1 5 5 0 0 5 5 2 3 2 0 3 3 1 1 3 0 1 4 1",
		"3 3 3 1 4 3 2 6 2 0 3 6 4 3 1 1 0 5 5 5 6 4 5 2 0 5 2 2 4 6 5 4 1 6 0 1 2 0 4 0 6 1",
		"3 3 5 3 0 4 0 2 2 5 1 0 2 6 4 5 5 5 0 5 1 3 3 1 0 0 2 3 1 2 6 6 4 4 2 1 4 4 1 6 6 6",
		"3 3 5 4 1 3 2 2 2 0 6 2 4 4 4 3 1 5 1 5 2 6 5 0 6 4 2 6 3 1 3 1 6 5 0 4 5 1 0 0 6 0",
		"3 5 5 1 1 3 3 2 6 4 5 6 5 1 1 3 4 3 2 5 0 6 3 5 4 4 0 0 1 0 1 6 0 0 4 2 4 2 6 2 6 2",
		"3 5 5 3 4 5 4 3 4 4 6 4 1 3 5 2 1 6 5 1 6 6 5 0 3 6 1 0 0 3 2 1 6 1 2 0 4 2 2 2 0 0",
		"3 6 5 1 0 3 5 2 6 3 4 0 1 5 0 4 4 3 6 6 6 0 3 3 2 2 2 0 6 1 0 5 2 2 1 1 5 1 4 4 4 5",
		"3 6 5 3 1 1 0 0 4 4 0 3 4 1 6 4 5 2 1 4 1 3 4 5 3 3 5 1 2 6 2 5 2 2 5 0 0 0 6 6 6 2",
		"3 6 5 4 0 5 0 0 1 1 6 0 1 1 5 6 5 4 4 4 3 6 0 2 1 6 1 3 3 2 5 3 3 0 4 4 2 5 2 2 2 6",
		"4 0 4 5 1 5 4 2 0 4 5 2 6 5 5 5 1 1 6 3 2 6 6 4 4 0 2 3 0 1 1 6 6 2 1 3 2 0 3 3 0 3",
		"4 1 2 3 1 6 0 4 6 3 6 3 5 0 4 5 0 1 5 5 1 4 3 1 3 5 2 0 4 3 4 2 6 0 2 1 2 6 5 2 0 6",
		"4 1 3 4 1 1 2 4 3 0 3 2 6 2 0 5 5 5 1 5 4 0 1 1 2 0 2 3 0 2 6 6 6 5 4 0 4 3 3 6 6 5",
		"4 2 0 5 1 0 1 3 4 3 5 4 3 6 0 2 4 2 0 1 0 3 6 0 6 3 1 5 5 5 2 6 5 3 4 6 2 2 4 1 6 1",
		"4 2 2 1 5 4 4 1 1 6 0 0 5 2 4 6 0 0 3 0 2 5 3 1 1 4 0 5 4 1 5 6 6 3 3 5 6 3 6 2 2 3",
		"4 2 2 3 0 5 3 4 5 0 5 5 6 0 3 1 5 4 6 1 0 6 3 5 1 4 0 0 1 3 1 1 4 4 3 2 6 2 6 2 6 2",
		"4 3 2 2 1 3 6 2 4 6 1 4 2 5 5 1 4 6 4 2 3 6 4 3 5 2 5 5 3 3 6 0 5 0 0 1 0 1 0 0 6 1",
		"4 3 2 2 3 1 0 0 6 5 4 5 6 5 3 3 1 3 5 2 1 4 2 3 0 6 6 1 6 6 2 1 1 2 4 0 4 4 0 0 5 5",
		"4 3 6 1 4 6 2 3 1 5 5 1 3 3 3 3 1 4 4 0 5 0 6 0 1 2 2 4 6 2 0 4 6 2 1 2 0 6 0 5 5 5",
		"4 3 6 6 1 2 1 6 6 5 6 3 6 1 1 1 5 1 0 0 4 4 2 2 3 4 2 5 5 4 4 5 0 2 5 2 3 0 3 3 0 0",
		"4 4 0 0 0 3 1 1 1 1 2 1 4 4 2 3 1 0 5 2 0 0 4 3 3 6 5 5 2 4 3 2 6 3 5 5 6 2 6 6 5 6",
		"4 5 0 1 4 4 6 2 5 3 3 2 4 6 4 2 5 5 2 6 4 3 5 1 1 1 2 2 3 3 6 3 1 6 6 0 5 1 0 0 0 0",
		"4 5 0 2 2 3 4 0 6 0 0 6 0 2 2 1 2 3 4 4 1 0 2 5 4 6 1 3 5 1 3 3 5 3 4 5 1 1 6 5 6 6",
		"4 5 1 1 0 3 2 4 6 5 0 4 3 0 1 4 3 6 2 2 5 5 4 2 1 4 2 5 3 0 2 1 6 1 6 0 5 3 6 6 3 0",
		"4 5 4 4 4 1 3 0 1 1 0 3 0 2 2 0 6 1 0 6 3 3 4 2 0 5 3 5 5 4 6 5 2 1 6 3 1 5 2 6 2 6",
		"4 6 2 4 2 2 6 3 5 6 5 0 3 1 1 6 5 6 3 0 0 4 3 5 0 5 4 2 1 3 0 0 5 3 4 4 6 1 1 2 1 2",
		"4 6 4 3 2 0 3 6 1 0 6 6 1 6 6 4 1 1 4 4 5 2 0 4 2 1 2 1 5 0 0 5 0 3 5 3 5 5 3 3 2 2",
		"4 6 5 3 1 0 5 3 5 6 2 1 1 0 6 0 4 4 6 5 6 1 1 3 2 1 3 2 5 2 0 4 0 6 4 4 5 3 0 3 2 2",
		"5 0 5 5 2 2 5 5 3 4 0 3 2 2 0 6 1 6 2 2 1 4 6 6 5 4 0 0 3 1 1 3 3 3 4 1 6 4 4 1 6 0",
		"5 1 4 3 5 5 5 2 1 3 6 2 4 4 0 4 2 2 6 6 5 5 2 2 3 0 6 3 6 1 1 4 3 0 4 1 6 1 0 3 0 0",
		"5 2 1 2 2 1 1 2 1 3 6 1 3 4 1 6 0 5 2 3 0 0 5 3 5 5 0 2 4 6 6 4 4 6 3 3 6 4 4 0 5 0",
		"5 2 3 2 4 4 5 6 3 2 3 5 2 0 0 2 2 0 6 3 1 3 0 3 0 5 5 0 6 5 6 6 1 1 1 1 4 1 4 6 4 4",
		"5 2 3 5 0 5 5 6 0 6 1 4 6 6 5 5 1 6 2 0 6 1 0 3 3 0 4 0 4 4 1 2 1 2 3 2 3 4 4 3 1 2",
		"5 3 0 1 1 5 1 3 5 6 1 5 4 2 2 3 5 1 3 0 4 2 4 3 6 0 6 4 5 4 6 3 2 6 4 1 2 2 6 0 0 0",
		"5 3 0 2 2 0 5 0 6 2 4 1 2 2 3 1 6 2 3 6 6 4 6 4 1 1 4 5 6 1 0 5 0 3 4 4 3 0 3 5 5 1",
		"5 3 1 3 2 4 2 5 4 6 4 4 2 0 1 4 5 2 2 3 3 1 6 0 0 4 3 5 6 6 0 3 2 1 1 0 6 5 5 1 0 6",
		"5 3 2 4 1 6 3 6 6 1 2 3 3 3 6 0 1 3 6 6 0 5 4 2 5 2 2 1 2 4 1 4 0 0 0 0 1 5 5 4 5 4",
		"5 3 3 3 1 2 5 6 3 2 4 6 5 3 2 1 6 6 1 1 4 2 6 2 0 5 3 2 5 5 0 4 4 1 4 6 1 0 4 0 0 0",
		"5 3 5 1 4 3 3 0 1 5 2 4 3 2 0 4 6 2 2 5 6 0 1 4 0 2 2 1 4 6 0 4 6 0 1 3 5 5 6 1 6 3",
		"5 4 0 6 4 0 1 2 3 5 6 1 3 2 2 6 0 6 2 0 5 3 2 3 0 5 4 2 1 0 4 6 5 1 1 3 3 4 5 1 4 6",
		"5 4 3 4 6 1 0 2 5 2 5 5 3 5 3 4 1 6 0 6 2 6 1 1 6 3 6 3 2 0 2 5 4 1 0 3 0 0 4 2 1 4",
		"5 4 3 6 0 3 1 3 6 6 1 6 0 6 6 3 4 1 4 0 3 2 1 1 4 4 4 2 2 2 2 2 0 1 5 5 0 5 0 5 3 5",
		"5 4 4 0 1 5 2 5 0 3 2 1 3 3 4 6 4 6 3 2 3 5 0 2 5 3 2 2 6 1 0 6 6 6 5 0 0 4 1 4 1 1",
		"5 5 3 6 6 4 0 1 0 5 6 1 2 6 1 6 6 0 4 1 5 3 2 0 1 3 2 2 3 1 3 4 2 2 4 3 5 4 5 0 4 0",
		"5 5 4 5 4 5 0 3 4 1 0 4 4 0 2 6 0 1 1 1 5 4 2 1 6 0 6 5 2 2 3 3 0 6 6 6 2 1 3 3 3 2",
		"5 5 6 2 1 5 0 3 4 5 6 6 4 6 5 5 6 0 4 3 2 4 1 2 1 6 0 0 0 4 0 1 4 3 3 2 3 2 3 1 1 2",
		"5 6 5 2 4 6 0 5 2 2 4 6 4 4 4 2 4 3 0 5 5 0 5 1 1 2 1 0 6 6 0 6 2 1 1 3 0 3 3 1 3 3",
		"5 6 5 3 0 4 1 6 5 1 1 6 3 5 1 1 2 0 4 2 6 5 2 2 3 0 5 3 3 1 0 3 2 0 6 2 0 4 4 6 4 4",
		"5 6 6 2 6 6 4 1 6 1 2 4 6 5 0 5 1 4 2 5 3 4 3 1 0 3 0 0 3 3 4 3 1 2 1 4 0 2 5 5 0 2",
		"6 0 0 6 2 3 2 4 3 3 1 1 0 4 0 0 0 6 3 2 1 6 5 3 2 3 6 4 5 1 2 1 5 1 6 2 4 5 5 5 4 4",
		"6 0 1 0 3 5 0 0 1 4 5 4 3 5 6 4 4 2 6 0 2 6 4 6 6 5 5 4 2 5 1 3 2 3 0 2 3 3 2 1 1 1",
		"6 0 2 0 2 1 1 4 5 2 5 1 3 4 4 4 4 5 1 3 1 4 3 1 5 3 5 0 6 2 3 6 0 2 3 6 0 5 0 6 6 2",
		"6 0 3 5 0 2 2 3 3 4 1 6 3 2 3 3 4 5 2 5 5 5 5 1 2 6 0 4 6 2 6 4 6 1 1 1 4 0 4 1 0 0",
		"6 0 4 2 6 5 1 0 2 6 0 1 3 0 3 0 0 5 5 3 3 3 6 6 6 1 5 2 2 5 2 4 1 2 5 3 4 4 4 4 1 1",
		"6 0 5 1 1 3 2 3 5 5 3 3 5 4 6 5 5 3 2 6 6 1 0 6 1 4 2 1 0 3 1 4 6 2 2 2 4 0 4 0 0 4",
		"6 0 5 3 0 3 5 1 4 0 0 5 1 0 2 4 5 5 5 2 1 0 6 6 1 2 3 1 6 6 6 1 4 3 2 4 4 4 2 3 3 2",
		"6 0 5 4 4 5 0 0 2 3 5 4 0 0 4 5 0 5 5 2 4 4 3 2 2 6 1 6 1 3 3 6 1 1 1 3 6 1 6 3 2 2",
		"6 0 6 2 5 6 5 3 3 2 6 0 0 3 4 4 1 0 2 4 5 3 4 5 1 3 5 6 4 4 0 6 2 2 3 2 0 5 1 1 1 1",
		"6 1 2 5 3 1 4 5 1 6 0 5 4 2 5 3 2 5 0 4 1 5 6 6 3 6 3 2 2 1 3 6 2 0 1 3 0 0 4 0 4 4",
		"6 1 3 6 2 4 6 3 4 6 6 4 6 0 4 0 3 3 4 5 1 0 3 3 5 5 2 1 1 1 0 5 0 5 0 4 5 1 2 2 2 2",
		"6 1 4 4 2 0 4 6 4 4 0 3 5 4 0 3 5 2 1 2 2 1 3 5 1 5 2 5 2 6 0 0 0 1 5 1 6 3 6 6 3 3",
		"6 2 1 3 2 5 0 1 3 3 6 3 3 6 5 1 6 1 5 5 4 6 2 0 2 4 0 6 1 0 5 2 1 5 0 4 0 3 2 4 4 4",
		"6 3 4 1 6 3 5 2 4 3 0 6 4 5 2 1 6 1 1 2 6 5 1 4 0 6 3 1 5 4 0 4 5 0 2 2 0 3 5 3 0 2",
		"6 4 2 2 2 6 5 3 4 2 0 4 6 3 5 5 5 3 0 0 2 6 0 5 1 1 0 2 3 1 6 1 1 5 1 3 4 3 0 6 4 4",
		"6 4 4 3 5 0 2 0 6 5 2 3 6 4 1 3 0 1 2 6 5 4 1 5 1 2 4 5 0 5 3 0 1 2 0 1 3 4 2 3 6 6",
		"6 5 0 3 1 3 3 1 2 6 4 1 0 3 1 6 4 2 1 1 2 5 6 0 0 0 0 4 5 2 5 3 3 4 4 5 2 6 6 5 2 4",
		"6 5 2 2 1 6 5 3 5 2 6 5 5 2 3 5 1 0 0 1 0 4 4 6 2 1 0 4 3 0 4 6 3 6 2 3 3 1 0 4 4 1",
		"6 5 2 6 5 3 0 4 2 1 4 4 1 0 0 6 4 5 4 2 2 2 0 2 6 0 0 3 1 4 1 1 3 3 3 5 3 5 6 6 1 5",
		"6 5 3 6 6 0 6 2 3 0 6 6 0 2 5 4 1 1 2 5 5 3 4 1 5 1 2 3 0 3 1 0 5 4 4 2 0 2 3 4 4 1",
		"6 6 3 0 2 3 2 2 2 6 6 5 5 1 6 4 3 2 0 2 5 3 6 0 1 1 5 5 3 4 5 4 4 0 0 3 1 4 4 0 1 1",
	})

	var tests []test

	for _, moves := range draws {
		tests = append(tests, test{moves, "Draw"})
	}

	for _, moves := range sweats {
		tests = append(tests, test{moves, "Yellow"})
	}

	const argc = 1000 // Preserve original argc

	for range argc*2 - len(tests) {
		moves, expected := emulPlay()

		tests = append(tests, test{
			strings.Trim(fmt.Sprint(moves), "[]"),
			expected,
		})
	}

	shuffle(tests)

	return outputTests(tests[:argc], tests[argc:])
})
