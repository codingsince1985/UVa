// UVa 10880 - Colin and Ryan

package main

import (
	"fmt"
	"os"
	"sort"
)

func solve(c, r int) []int {
	if c == r {
		return []int{0}
	}
	var q []int
	for i := 1; i*i <= c-r; i++ {
		if (c-r)%i == 0 {
			if i > r {
				q = append(q, i)
			}
			if f := (c - r) / i; f > r && f != i {
				q = append(q, f)
			}
		}
	}
	sort.Ints(q)
	return q
}

func main() {
	in, _ := os.Open("10880.in")
	defer in.Close()
	out, _ := os.Create("10880.out")
	defer out.Close()

	var n, c, r int
	fmt.Fscanf(in, "%d", &n)
	for kase := 1; kase <= n; kase++ {
		fmt.Fscanf(in, "%d%d", &c, &r)
		fmt.Fprintf(out, "Case #%d:", kase)
		if q := solve(c, r); len(q) > 0 {
			s := fmt.Sprintf("%v", q)
			fmt.Fprintf(out, " %s", s[1:len(s)-1])
		}
		fmt.Fprintln(out)
	}
}
