// UVa 11000 - Bee

package main

import (
	"fmt"
	"os"
)

func solve(n int) (uint64, uint64) {
	var f, m uint64 = 1, 0
	for ; n > 0; n-- {
		m, f = f+m, m+1
	}
	return m, f
}

func main() {
	in, _ := os.Open("11000.in")
	defer in.Close()
	out, _ := os.Create("11000.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == -1 {
			break
		}
		m, f := solve(n)
		fmt.Fprintf(out, "%d %d\n", m, m+f)
	}
}
