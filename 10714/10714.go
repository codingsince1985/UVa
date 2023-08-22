// UVa 10714 - Ants

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10714.in")
	defer in.Close()
	out, _ := os.Create("10714.out")
	defer out.Close()

	var kase, l, n, pos, longest, shortest int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		for fmt.Fscanf(in, "%d%d", &l, &n); n > 0; n-- {
			fmt.Fscanf(in, "%d", &pos)
			shortest = max(shortest, min(pos, l-pos))
			longest = max(longest, max(pos, l-pos))
		}
		fmt.Fprintln(out, shortest, longest)
	}
}
