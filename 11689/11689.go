// UVa 11689 - Soda Surpler

package main

import (
	"fmt"
	"os"
)

func solve(e, f, c int) int {
	var total int
	bottles := e + f
	for bottles >= c {
		newBottles := bottles / c
		total += newBottles
		bottles %= c
		bottles += newBottles
	}
	return total
}

func main() {
	in, _ := os.Open("11689.in")
	defer in.Close()
	out, _ := os.Create("11689.out")
	defer out.Close()

	var kase, e, f, c int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d%d%d", &e, &f, &c)
		fmt.Fprintln(out, solve(e, f, c))
	}
}
