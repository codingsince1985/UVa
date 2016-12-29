// UVa 10714 - Ants

package main

import (
	"fmt"
	"os"
)

func sort(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func main() {
	in, _ := os.Open("10714.in")
	defer in.Close()
	out, _ := os.Create("10714.out")
	defer out.Close()

	var kase, l, n, pos, max, min int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		for fmt.Fscanf(in, "%d%d", &l, &n); n > 0; n-- {
			fmt.Fscanf(in, "%d", &pos)
			s, b := sort(pos, l-pos)
			if s > min {
				min = s
			}
			if b > max {
				max = b
			}
		}
		fmt.Fprintln(out, min, max)
	}
}
