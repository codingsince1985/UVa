// UVa 11608 - No Problem

package main

import (
	"fmt"
	"os"
)

var out *os.File

func solve(s int, p, c []int) {
	for i := range p {
		if s >= c[i] {
			fmt.Fprintln(out, "No problem! :D")
			s += (p[i] - c[i])
		} else {
			fmt.Fprintln(out, "No problem. :(")
			s += p[i]
		}
	}
}

func main() {
	in, _ := os.Open("11608.in")
	defer in.Close()
	out, _ = os.Create("11608.out")
	defer out.Close()

	var s int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &s); s < 0 {
			break
		}
		p := make([]int, 12)
		for i := range p {
			fmt.Fscanf(in, "%d", &p[i])
		}
		c := make([]int, 12)
		for i := range c {
			fmt.Fscanf(in, "%d", &c[i])
		}
		fmt.Fprintf(out, "Case %d:\n", kase)
		solve(s, p, c)
	}
}
