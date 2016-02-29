// UVa 10360 - Rat Attack

package main

import (
	"fmt"
	"os"
)

const MAX = 1025

func process(d, x, y, population int, rats *[MAX][MAX]int) {
	for i := x - d; i <= x+d; i++ {
		if i < 0 || i >= MAX {
			continue
		}
		for j := y - d; j <= y+d; j++ {
			if j < 0 || j >= MAX {
				continue
			}
			rats[i][j] += population
		}
	}
}

func output(out *os.File, rats *[MAX][MAX]int) {
	max, x, y := 0, 0, 0
	for i := range rats {
		for j := range rats[i] {
			if rats[i][j] > max {
				max = rats[i][j]
				x, y = i, j
			}
		}
	}
	fmt.Fprintln(out, x, y, max)
}

func main() {
	in, _ := os.Open("10360.in")
	defer in.Close()
	out, _ := os.Create("10360.out")
	defer out.Close()

	var kase, d, n, x, y, population int
	fmt.Fscanf(in, "%d", &kase)
	for i := 0; i < kase; i++ {
		var rats [MAX][MAX]int
		fmt.Fscanf(in, "%d\n%d", &d, &n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(in, "%d%d%d", &x, &y, &population)
			process(d, x, y, population, &rats)
		}
		output(out, &rats)
	}
}
