// UVa 10189 - Minesweeper

package main

import (
	"fmt"
	"os"
)

var (
	m, n  int
	delta = []int{-1, 0, 1}
)

func solve(field [][]byte) [][]byte {
	res := make([][]byte, m)
	for i := range field {
		res[i] = make([]byte, n)
		for j := range field[i] {
			if field[i][j] == '*' {
				res[i][j] = '*'
			} else {
				res[i][j] = '0'
				for _, dx := range delta {
					for _, dy := range delta {
						x, y := i+dx, j+dy
						if !(dx == 0 && dy == 0) && x >= 0 && y >= 0 && x < m && y < n && field[x][y] == '*' {
							res[i][j]++
						}
					}
				}
			}
		}
	}
	return res
}

func output(out *os.File, field [][]byte) {
	for _, v := range field {
		fmt.Fprintln(out, string(v))
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("10189.in")
	defer in.Close()
	out, _ := os.Create("10189.out")
	defer out.Close()

	var kase int
	var line string
	for {
		if fmt.Fscanf(in, "%d%d", &m, &n); m == 0 {
			break
		}
		kase++
		fmt.Fprintf(out, "Field #%d:\n", kase)
		field := make([][]byte, m)
		for i := range field {
			fmt.Fscanf(in, "%s", &line)
			field[i] = []byte(line)
		}
		field = solve(field)
		output(out, field)
	}
}
