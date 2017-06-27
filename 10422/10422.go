// UVa 10422 - Knights in FEN

package main

import (
	"fmt"
	"os"
)

const max = 11

var (
	steps int
	board [5][5]byte
	final = [5][5]byte{
		{'1', '1', '1', '1', '1'},
		{'0', '1', '1', '1', '1'},
		{'0', '0', ' ', '1', '1'},
		{'0', '0', '0', '0', '1'},
		{'0', '0', '0', '0', '0'},
	}
	directions = [][2]int{{-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}}
)

func dfs(y, x, level int) {
	if level >= steps {
		return
	}
	if final == board {
		steps = level
		return
	}
	for _, direction := range directions {
		if newY, newX := y+direction[0], x+direction[1]; newY >= 0 && newY < 5 && newX >= 0 && newX < 5 {
			board[y][x], board[newY][newX] = board[newY][newX], board[y][x]
			dfs(newY, newX, level+1)
			board[y][x], board[newY][newX] = board[newY][newX], board[y][x]
		}
	}
}

func main() {
	in, _ := os.Open("10422.in")
	defer in.Close()
	out, _ := os.Create("10422.out")
	defer out.Close()

	var y, x, kase int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		for i, row := range board {
			for j := range row {
				if fmt.Fscanf(in, "%c", &board[i][j]); board[i][j] == ' ' {
					y, x = i, j
				}
			}
			fmt.Fscanln(in)
		}
		steps = max
		dfs(y, x, 0)
		if steps == max {
			fmt.Fprintln(out, "Unsolvable in less than 11 move(s).")
		} else {
			fmt.Fprintf(out, "Solvable in %d move(s).\n", steps)
		}
	}
}
