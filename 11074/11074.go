// UVa 11074 - Draw Grid

package main

import (
	"fmt"
	"os"
	"strings"
)

func solve(s, t, n int) string {
	w := t + (s+t)*n
	return strings.Repeat(strings.Repeat(strings.Repeat("*", w)+"\n", t)+strings.Repeat(strings.Repeat(strings.Repeat("*", t)+strings.Repeat(".", s), n)+strings.Repeat("*", t)+"\n", s), n) + strings.Repeat(strings.Repeat("*", w)+"\n", t)
}

func main() {
	in, _ := os.Open("11074.in")
	defer in.Close()
	out, _ := os.Create("11074.out")
	defer out.Close()

	var s, t, n int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d%d%d", &s, &t, &n); n == 0 {
			break
		}
		fmt.Fprintf(out, "Case %d:\n%s\n", kase, solve(s, t, n))
	}
}
