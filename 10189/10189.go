// UVa 10189 - Minesweeper

package main

import (
	"fmt"
	"os"
)

var (
	m, n int
	dx   = []int{-1, 0, 1}
	dy   = []int{-1, 0, 1}
)

func solve(field [][]byte) [][]byte {
	res := make([][]byte, m)
	for i := range field {
		res[i] = make([]byte, n)
		for j := range field[i] {
			if field[i][j] == '*' {
				res[i][j] = '*'
				continue
			}
			var cnt byte = 0
			for _, k := range dx {
				for _, l := range dy {
					newX := i + k
					newY := j + l
					if (k == 0 && l == 0) || newX < 0 || newY < 0 || newX >= m || newY >= n {
						continue
					}
					if field[newX][newY] == '*' {
						cnt++
					}
				}
			}
			res[i][j] = '0' + cnt
		}
	}
	return res
}

func output(out *os.File, field [][]byte) {
	for i := range field {
		fmt.Fprintln(out, string(field[i]))
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("10189.in")
	defer in.Close()
	out, _ := os.Create("10189.out")
	defer out.Close()

	var kase int
	var field [][]byte
	var line string
	for {
		if fmt.Fscanf(in, "%d%d", &m, &n); m == 0 {
			break
		}
		kase++
		fmt.Fprintf(out, "Field #%d:\n", kase)
		field = make([][]byte, m)
		for i := range field {
			fmt.Fscanf(in, "%s", &line)
			field[i] = []byte(line)
		}
		field = solve(field)
		output(out, field)
	}
}
