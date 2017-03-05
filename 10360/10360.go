// UVa 10360 - Rat Attack

package main

import (
	"fmt"
	"os"
)

const max = 1025

func process(d, x, y, population int, rats *[max][max]int) {
	for i := x - d; i <= x+d; i++ {
		if i < 0 || i >= max {
			continue
		}
		for j := y - d; j <= y+d; j++ {
			if j < 0 || j >= max {
				continue
			}
			rats[i][j] += population
		}
	}
}

func output(out *os.File, rats *[max][max]int) {
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
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		var rats [max][max]int
		for fmt.Fscanf(in, "%d\n%d", &d, &n); n > 0; n-- {
			fmt.Fscanf(in, "%d%d%d", &x, &y, &population)
			process(d, x, y, population, &rats)
		}
		output(out, &rats)
	}
}
