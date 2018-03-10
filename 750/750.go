// UVa 750 - 8 Queens Chess Problem

package main

import (
	"fmt"
	"io"
	"os"
)

var (
	row, count int
	out        io.WriteCloser
)

func in(x, y int) bool { return x >= 0 && x < 8 && y >= 0 && y < 8 }

func valid(chess [8][8]bool, r, c int) bool {
	for i := 0; i < 8; i++ {
		if chess[i][c] {
			return false
		}
	}
	for x, y := r, c; in(x, y); x, y = x+1, y-1 {
		if chess[x][y] {
			return false
		}
	}
	for x, y := r, c; in(x, y); x, y = x-1, y+1 {
		if chess[x][y] {
			return false
		}
	}
	for x, y := r, c; in(x, y); x, y = x-1, y-1 {
		if chess[x][y] {
			return false
		}
	}
	for x, y := r, c; in(x, y); x, y = x+1, y+1 {
		if chess[x][y] {
			return false
		}
	}
	return true
}

func backtracking(r, cnt int, chess [8][8]bool) {
	if cnt == 8 {
		count++
		fmt.Fprintf(out, "%2d     ", count)
		for i := range chess {
			for j := range chess[i] {
				if chess[i][j] {
					fmt.Fprint(out, " ", j+1)
				}
			}
		}
		fmt.Fprintln(out)
		return
	}
	if r == row-1 && row != 8 {
		backtracking(r+1, cnt, chess)
	} else {
		for c := 0; c < 8; c++ {
			if valid(chess, r, c) {
				chess[r][c] = true
				backtracking(r+1, cnt+1, chess)
				chess[r][c] = false
			}
		}
	}
}

func main() {
	in, _ := os.Open("750.in")
	defer in.Close()
	out, _ = os.Create("750.out")
	defer out.Close()

	var c int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &row, &c); err != nil {
			break
		}
		fmt.Fprintln(out, "SOLN       COLUMN")
		fmt.Fprintln(out, " #      1 2 3 4 5 6 7 8")
		fmt.Fprintln(out)
		var chess [8][8]bool
		chess[row-1][c-1] = true
		backtracking(0, 1, chess)
	}
}
