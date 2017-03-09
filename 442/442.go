// UVa 442 - Matrix Chain Multiplication

package main

import (
	"fmt"
	"os"
)

var m = make(map[string][2]int)

func solve(out *os.File, s string) {
	if len(s) == 1 {
		fmt.Fprintln(out, 0)
		return
	}

	var e [][2]int
	var count int
	var a, b [2]int
	for i := range s {
		switch s[i] {
		case '(':
		case ')':
			b = e[len(e)-1]
			a = e[len(e)-2]
			if a[1] != b[0] {
				fmt.Fprintln(out, "error")
				return
			}
			count += a[0] * a[1] * b[1]
			e = e[:len(e)-2]
			e = append(e, [2]int{a[0], b[1]})
		default:
			e = append(e, m[string(s[i])])
		}
	}
	fmt.Fprintln(out, count)
}

func main() {
	in, _ := os.Open("442.in")
	defer in.Close()
	out, _ := os.Create("442.out")
	defer out.Close()

	var n, r, c int
	var s string
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%s%d%d", &s, &r, &c)
		m[s] = [2]int{r, c}
	}
	for {
		if _, err := fmt.Fscanf(in, "%s", &s); err != nil {
			break
		}
		solve(out, s)
	}
}
