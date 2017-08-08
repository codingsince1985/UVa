// UVa 10300 - Ecological Premium

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10300.in")
	defer in.Close()
	out, _ := os.Create("10300.out")
	defer out.Close()

	var n, f int
	var s, a, e uint64
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		var t uint64
		for fmt.Fscanf(in, "%d", &f); f > 0; f-- {
			fmt.Fscanf(in, "%d%d%d", &s, &a, &e)
			t += s * e
		}
		fmt.Fprintln(out, t)
	}
}
