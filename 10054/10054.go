// UVa 10054 - The Necklace

package main

import (
	"fmt"
	"io"
	"os"
)

const max = 50

var (
	out    io.WriteCloser
	matrix [][]int
	degree []int
)

func euler(v int) {
	for i := range matrix {
		if matrix[v][i] != 0 {
			matrix[v][i]--
			matrix[i][v]--
			euler(i)
			fmt.Fprintln(out, i, v)
		}
	}
}

func solve(v int) {
	for _, vi := range degree {
		if vi%2 != 0 {
			fmt.Fprintln(out, "some beads may be lost")
			return
		}
	}
	euler(v)
}

func main() {
	in, _ := os.Open("10054.in")
	defer in.Close()
	out, _ = os.Create("10054.out")
	defer out.Close()

	var t, n, l, r int
	fmt.Fscanf(in, "%d", &t)
	for kase := 1; kase <= t; kase++ {
		matrix = make([][]int, max+1)
		for i := range matrix {
			matrix[i] = make([]int, max+1)
		}
		degree = make([]int, max+1)
		for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
			fmt.Fscanf(in, "%d%d", &l, &r)
			matrix[l][r]++
			matrix[r][l]++
			degree[l]++
			degree[r]++
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
		fmt.Fprintf(out, "Case #%d\n", kase)
		solve(r)
	}
}
