// UVa 712 - S-Trees

package main

import (
	"fmt"
	"os"
)

func solve(path string) int {
	pointer := 0
	for _, child := range []byte(path) {
		if pointer *= 2; child == '1' {
			pointer++
		}
	}
	return pointer
}

func main() {
	in, _ := os.Open("712.in")
	defer in.Close()
	out, _ := os.Create("712.out")
	defer out.Close()

	var n, m int
	var x string
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			fmt.Fscanf(in, "%s", &x)
		}
		fmt.Fscanf(in, "%s", &x)
		leaves := []byte(x)
		fmt.Fprintf(out, "S-Tree #%d:\n", kase)
		for fmt.Fscanf(in, "%d", &m); m > 0; m-- {
			fmt.Fscanf(in, "%s", &x)
			fmt.Fprintf(out, "%c", leaves[solve(x)])
		}
		fmt.Fprintln(out, "\n")
	}
}
