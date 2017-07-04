// UVa 10493 - Cats, with or without Hats

package main

import (
	"fmt"
	"os"
)

func solve(out *os.File, n, m int) {
	fmt.Fprintf(out, "%d %d ", n, m)
	switch {
	case n == 1 && m == 1:
		fmt.Fprintln(out, "Multiple")
	case n == 1 || (n*m-1)%(n-1) != 0:
		fmt.Fprintln(out, "Impossible")
	default:
		fmt.Fprintln(out, (n*m-1)/(n-1))
	}
}

func main() {
	in, _ := os.Open("10493.in")
	defer in.Close()
	out, _ := os.Create("10493.out")
	defer out.Close()

	var n, m int
	for {
		if fmt.Fscanf(in, "%d%d", &n, &m); n == 0 {
			break
		}
		solve(out, n, m)
	}
}
