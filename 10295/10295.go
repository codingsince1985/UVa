// UVa 10295 - Hay Points

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10295.in")
	defer in.Close()
	out, _ := os.Create("10295.out")
	defer out.Close()

	var m, n, d int
	var w string
	dict := make(map[string]int)
	fmt.Fscanf(in, "%d%d", &m, &n)
	for ; m > 0; m-- {
		fmt.Fscanf(in, "%s%d", &w, &d)
		dict[w] = d
	}

	for ; n > 0; n-- {
		var s uint64 = 0
		for {
			if fmt.Fscanf(in, "%s", &w); w == "." {
				break
			}
			s += uint64(dict[w])
		}
		fmt.Fprintln(out, s)
	}
}
