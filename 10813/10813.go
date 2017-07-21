// UVa 10813 - Traditional BINGO

package main

import (
	"fmt"
	"os"
)

func done(board [][]int) bool {
	var sum int
	for i := range board {
		for j := range board {
			sum += board[i][j]
		}
		if sum == 0 {
			return true
		}
		sum = 0
		for j := range board {
			sum += board[j][i]
		}
		if sum == 0 {
			return true
		}
	}
	sum = 0
	for i := range board {
		sum += board[i][i]
	}
	if sum == 0 {
		return true
	}
	sum = 0
	for i := range board {
		sum += board[i][4-i]
	}
	return sum == 0
}

func solve(board [][]int, b int) bool {
	for i, row := range board {
		for j, cell := range row {
			if cell == b {
				board[i][j] = 0
			}
		}
	}
	return done(board)
}

func main() {
	in, _ := os.Open("10813.in")
	defer in.Close()
	out, _ := os.Create("10813.out")
	defer out.Close()

	var kase, b int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		board := make([][]int, 5)
		for i := range board {
			board[i] = make([]int, 5)
			for j := range board[i] {
				if i != 2 || j != 2 {
					fmt.Fscanf(in, "%d", &board[i][j])
				}
			}
		}
		for i := 1; i <= 75; i++ {
			fmt.Fscanf(in, "%d", &b)
			if solve(board, b) {
				fmt.Fprintf(out, "BINGO after %d numbers announced\n", i)
				break
			}
		}
	}
}
