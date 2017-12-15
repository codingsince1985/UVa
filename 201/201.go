// UVa 201 - Squares

package main

import (
	"fmt"
	"os"
)

type point struct{ h, v bool }

var n int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isSquare(i, j, k int, grid [][]point) bool {
	for d := 0; d < k; d++ {
		if !grid[i+d][j].v || !grid[i+d][j+k].v || !grid[i][j+d].h || !grid[i+k][j+d].h {
			return false
		}
	}
	return true
}

func solve(grid [][]point) map[int]int {
	squareMap := make(map[int]int)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			max := min(n-1-i, n-1-j)
			for k := 1; k <= max; k++ {
				if isSquare(i, j, k, grid) {
					squareMap[k]++
				}
			}
		}
	}
	return squareMap
}

func main() {
	in, _ := os.Open("201.in")
	defer in.Close()
	out, _ := os.Create("201.out")
	defer out.Close()

	var m, n1, n2 int
	var direction byte
	for kase := 1; ; kase++ {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		grid := make([][]point, n)
		for i := range grid {
			grid[i] = make([]point, n)
		}
		for fmt.Fscanf(in, "%d", &m); m > 0; m-- {
			switch fmt.Fscanf(in, "%c%d%d", &direction, &n1, &n2); direction {
			case 'H':
				grid[n1-1][n2-1].h = true
			case 'V':
				grid[n2-1][n1-1].v = true
			}
		}
		squareMap := solve(grid)
		if kase > 1 {
			fmt.Fprint(out, "\n**********************************\n\n")
		}
		if fmt.Fprintf(out, "Problem #%d\n\n", kase); len(squareMap) == 0 {
			fmt.Fprintln(out, "No completed squares can be found.")
		} else {
			for i := 1; i < n; i++ {
				if num, ok := squareMap[i]; ok {
					fmt.Fprintf(out, "%d square (s) of size %d\n", num, i)
				}
			}
		}
	}
}
