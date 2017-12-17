// UVa 10363 - Tic Tac Toe

package main

import (
	"fmt"
	"os"
)

var grid [][]byte

func win(player byte) bool {
	for _, row := range grid {
		if row[0] == row[1] && row[1] == row[2] && row[2] == player {
			return true
		}
	}
	for i := 0; i < 3; i++ {
		if grid[0][i] == grid[1][i] && grid[1][i] == grid[2][i] && grid[2][i] == player {
			return true
		}
	}
	return grid[0][0] == grid[1][1] && grid[1][1] == grid[2][2] && grid[2][2] == player ||
		grid[2][0] == grid[1][1] && grid[1][1] == grid[0][2] && grid[0][2] == player
}

func count(player byte) int {
	var num int
	for _, row := range grid {
		for _, cell := range row {
			if cell == player {
				num++
			}
		}
	}
	return num
}

func solve() bool {
	switch xs, os, xwin, owin := count('X'), count('O'), win('X'), win('O'); {
	case !xwin && !owin:
		return xs == os || xs-1 == os
	case xwin && !owin:
		return xs-1 == os
	case owin && !xwin:
		return xs == os
	default:
		return false
	}
}

func main() {
	in, _ := os.Open("10363.in")
	defer in.Close()
	out, _ := os.Create("10363.out")
	defer out.Close()

	var n int
	var line string
	fmt.Fscanf(in, "%d", &n)
	for kase := 1; kase <= n; kase++ {
		grid = make([][]byte, 3)
		for i := range grid {
			fmt.Fscanf(in, "%s", &line)
			grid[i] = []byte(line)
		}
		if solve() {
			fmt.Fprintln(out, "yes")
		} else {
			fmt.Fprintln(out, "no")
		}
		fmt.Fscanln(in)
	}
}
