// UVa 10667 - Largest Block

package main

import (
	"fmt"
	"os"
)

type cell struct {
	occupied bool
	up       int
}

func verticalSum(board [][]cell) {
	for i := range board {
		for j := range board[i] {
			if board[i][j].occupied {
				board[i][j].up = 0
			} else {
				if i == 0 {
					board[i][j].up = 1
				} else {
					board[i][j].up = board[i-1][j].up + 1
				}
			}
		}
	}
}

func solve(s int, board [][]cell) int {
	verticalSum(board)
	var maxArea int
	for i := range board {
		for j := range board[i] {
			sum := board[i][j].up
			for k := j + 1; k < s && board[i][k].up >= board[i][j].up; k++ {
				sum += board[i][j].up
			}
			for k := j - 1; k >= 0 && board[i][k].up >= board[i][j].up; k-- {
				sum += board[i][j].up
			}
			if sum > maxArea {
				maxArea = sum
			}
		}
	}
	return maxArea
}

func main() {
	in, _ := os.Open("10667.in")
	defer in.Close()
	out, _ := os.Create("10667.out")
	defer out.Close()

	var kase, s, b, r1, c1, r2, c2 int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d", &s)
		board := make([][]cell, s)
		for i := range board {
			board[i] = make([]cell, s)
		}
		for fmt.Fscanf(in, "%d", &b); b > 0; b-- {
			fmt.Fscanf(in, "%d%d%d%d", &r1, &c1, &r2, &c2)
			for i := r1 - 1; i < r2; i++ {
				for j := c1 - 1; j < c2; j++ {
					board[i][j].occupied = true
				}
			}
		}
		fmt.Fprintln(out, solve(s, board))
	}
}
