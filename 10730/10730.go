// UVa 10730 - Antiarithmetic?

package main

import (
	"fmt"
	"os"
)

func solve(n int, pos []int) bool {
	for i := range pos {
		for j := 1; i+2*j < n; j++ {
			if (pos[i] < pos[i+j] && pos[i+j] < pos[i+j*2]) ||
				(pos[i] > pos[i+j] && pos[i+j] > pos[i+j*2]) {
				return false
			}
		}
	}
	return true
}

func main() {
	in, _ := os.Open("10730.in")
	defer in.Close()
	out, _ := os.Create("10730.out")
	defer out.Close()

	var m, n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fscanf(in, "%c")
		pos := make([]int, n)
		for i := range pos {
			fmt.Fscanf(in, "%d", &m)
			pos[m] = i
		}
		if solve(n, pos) {
			fmt.Fprintln(out, "yes")
		} else {
			fmt.Fprintln(out, "no")
		}
	}
}
