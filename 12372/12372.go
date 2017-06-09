// UVa 12372 - Packing for Holiday

package main

import (
	"fmt"
	"os"
)

const size = 20

func main() {
	in, _ := os.Open("12372.in")
	defer in.Close()
	out, _ := os.Create("12372.out")
	defer out.Close()

	var kase, l, w, h int
	fmt.Fscanf(in, "%d", &kase)
	for i := 1; i <= kase; i++ {
		if fmt.Fscanf(in, "%d%d%d", &l, &w, &h); l <= size && w <= size && h <= size {
			fmt.Fprintf(out, "Case %d: good\n", i)
		} else {
			fmt.Fprintf(out, "Case %d: bad\n", i)
		}
	}
}
