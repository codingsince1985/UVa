// UVa 10849 - Move the bishop

package main

import (
	"fmt"
	"os"
)

type cell struct{ row, column int }

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func solve(bishop, target cell) string {
	dr, dc := abs(bishop.row-target.row), abs(bishop.column-target.column)
	switch {
	case dr == 0 && dc == 0:
		return "0"
	case dr == dc:
		return "1"
	case abs(dr-dc)%2 == 1:
		return "no move"
	default:
		return "2"
	}
}

func main() {
	in, _ := os.Open("10849.in")
	defer in.Close()
	out, _ := os.Create("10849.out")
	defer out.Close()

	var c, t, n int
	var bishop, target cell
	for fmt.Fscanf(in, "%d", &c); c > 0; c-- {
		for fmt.Fscanf(in, "\n%d\n%d", &t, &n); t > 0; t-- {
			fmt.Fscanf(in, "%d%d%d%d", &bishop.row, &bishop.column, &target.row, &target.column)
			fmt.Fprintln(out, solve(bishop, target))
		}
	}
}
