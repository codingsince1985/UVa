// UVa 10054 - The Necklace

package main

import (
	"fmt"
	"os"
)

var (
	out    *os.File
	matrix [51][51]int
	degree [51]int
)

func initialize() {
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}
	for i := range degree {
		degree[i] = 0
	}
}

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
		initialize()
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
