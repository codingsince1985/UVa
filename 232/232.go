// UVa 232 - Crossword Answers

package main

import (
	"fmt"
	"os"
)

const (
	across = 1 << iota
	down
)

var (
	r, c int
	out  *os.File
)

type cell struct{ r, c, v int }

func order(board [][]byte) []cell {
	var order []cell
	for i, row := range board {
		for j := range row {
			if row[j] != '*' {
				var value int
				if j == 0 || row[j-1] == '*' {
					value += across
				}
				if i == 0 || board[i-1][j] == '*' {
					value += down
				}
				if value > 0 {
					order = append(order, cell{i, j, value})
				}
			}
		}
	}
	return order
}

func outputAcross(board [][]byte, order []cell) {
	fmt.Fprintln(out, "Across")
	for i, curr := range order {
		if curr.v&across != 0 {
			fmt.Fprintf(out, "%3d.", i+1)
			for ptr := curr.c; ptr < c && board[curr.r][ptr] != '*'; ptr++ {
				fmt.Fprintf(out, "%c", board[curr.r][ptr])
			}
			fmt.Fprintln(out)
		}
	}
}

func outputDown(board [][]byte, order []cell) {
	fmt.Fprintln(out, "Down")
	for i, curr := range order {
		if curr.v&down != 0 {
			fmt.Fprintf(out, "%3d.", i+1)
			for ptr := curr.r; ptr < r && board[ptr][curr.c] != '*'; ptr++ {
				fmt.Fprintf(out, "%c", board[ptr][curr.c])
			}
			fmt.Fprintln(out)
		}
	}
}

func main() {
	in, _ := os.Open("232.in")
	defer in.Close()
	out, _ = os.Create("232.out")
	defer out.Close()

	var line string
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &r); r == 0 {
			break
		}
		fmt.Fscanf(in, "%d", &c)
		board := make([][]byte, r)
		for i := range board {
			fmt.Fscanf(in, "%s", &line)
			board[i] = []byte(line)
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
		fmt.Fprintf(out, "puzzle #%d:\n", kase)
		mark := order(board)
		outputAcross(board, mark)
		outputDown(board, mark)
	}
}
