// UVa 379 - Hi-Q

package main

import (
	"fmt"
	"os"
)

const size = 7

var (
	pos = [][]int{
		{0, 0, 1, 2, 3, 0, 0},
		{0, 0, 4, 5, 6, 0, 0},
		{7, 8, 9, 10, 11, 12, 13},
		{14, 15, 16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25, 26, 27},
		{0, 0, 28, 29, 30, 0, 0},
		{0, 0, 31, 32, 33, 0, 0},
	}
	directions = [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
)

func in(y, x int) bool { return y < size && y >= 0 && x < size && x >= 0 && pos[y][x] != 0 }

func solve(board [][]bool) int {
here:
	for tY := size - 1; tY >= 0; tY-- {
		for tX := size - 1; tX >= 0; tX-- {
			if pos[tY][tX] != 0 && !board[tY][tX] {
				for _, direction := range directions {
					if sY, sX := tY+2*direction[1], tX+2*direction[0]; in(sY, sX) && board[sY][sX] {
						if mY, mX := tY+direction[1], tX+direction[0]; in(mY, mX) && board[mY][mX] {
							board[tY][tX], board[mY][mX], board[sY][sX] = true, false, false
							goto here
						}
					}
				}
			}
		}
	}
	var total int
	for i, row := range board {
		for j, cell := range row {
			if cell {
				total += pos[i][j]
			}
		}
	}
	return total
}

func main() {
	in, _ := os.Open("379.in")
	defer in.Close()
	out, _ := os.Create("379.out")
	defer out.Close()

	fmt.Fprintln(out, "HI Q OUTPUT")
	var kase, n int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		board := make([][]bool, size)
		for i := range board {
			board[i] = make([]bool, size)
		}
		for {
			if fmt.Fscanf(in, "%d", &n); n == 0 {
				break
			}
			for i, row := range pos {
				for j, cell := range row {
					if cell == n {
						board[i][j] = true
						break
					}
				}
			}
		}
		fmt.Fprintln(out, solve(board))
	}
	fmt.Fprintln(out, "END OF OUTPUT")
}
