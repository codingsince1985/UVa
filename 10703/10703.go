// UVa 10703 - Free spots

package main

import (
	"fmt"
	"os"
)

func cover(board [][]bool, x1, y1, x2, y2 int) {
	if x2 < x1 {
		x1, x2 = x2, x1
	}
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	for i := x1 - 1; i < x2; i++ {
		for j := y1 - 1; j < y2; j++ {
			board[i][j] = true
		}
	}
}

func uncovered(board [][]bool) int {
	var count int
	for _, vi := range board {
		for _, vj := range vi {
			if !vj {
				count++
			}
		}
	}
	return count
}

func main() {
	in, _ := os.Open("10703.in")
	defer in.Close()
	out, _ := os.Create("10703.out")
	defer out.Close()

	var w, h, n, x1, y1, x2, y2 int
	for {
		if fmt.Fscanf(in, "%d%d%d", &w, &h, &n); w == 0 && h == 0 && n == 0 {
			break
		}
		board := make([][]bool, w)
		for i := range board {
			board[i] = make([]bool, h)
		}
		for n > 0 {
			fmt.Fscanf(in, "%d%d%d%d", &x1, &y1, &x2, &y2)
			cover(board, x1, y1, x2, y2)
			n--
		}
		fmt.Fscanln(in)
		count := uncovered(board)
		switch count {
		case 0:
			fmt.Fprintln(out, "There is no empty spots.")
		case 1:
			fmt.Fprintln(out, "There is one empty spot.")
		default:
			fmt.Fprintf(out, "There are %d empty spots.\n", count)
		}
	}
}
