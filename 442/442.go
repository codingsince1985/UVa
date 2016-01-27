// UVa 442 - Matrix Chain Multiplication

package main

import (
	"fmt"
	"os"
)

var m map[string][2]int

func solve(out *os.File, s string) {
	l := len(s)
	if l == 1 {
		fmt.Fprintln(out, 0)
		return
	}

	var e [][2]int
	var p []string
	count := 0
	var a, b [2]int

	for _, v := range s {
		c := string(v)
		if c == "(" {
			p = append(p, c)
		} else if c == ")" {
			b = e[len(e) - 1]
			a = e[len(e) - 2]
			if a[1] != b[0] {
				fmt.Fprintln(out, "error")
				return
			}
			count += a[0] * a[1] * b[1]

			e = e[:len(e) - 2]
			e = append(e, [2]int{a[0], b[1]})
			p = p[:len(p) - 1]
		} else {
			e = append(e, m[string(c)])
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
	fmt.Fscanf(in, "%d", &n)
	m = make(map[string][2]int)
	for i := 0; i < n; i++ {
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