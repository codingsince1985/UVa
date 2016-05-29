// UVa 10530 - Guessing Game

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10530.in")
	defer in.Close()
	out, _ := os.Create("10530.out")
	defer out.Close()

	var n int
	l, h, cheating := 1, 10, false
	var s1, s2 string
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fscanf(in, "%s%s", &s1, &s2)
		switch {
		case s2 == "high":
			h = n - 1
		case s2 == "low":
			l = n + 1
		default:
			if cheating || n < l || n > h {
				fmt.Fprintln(out, "Stan is dishonest")
			} else {
				fmt.Fprintln(out, "Stan may be honest")
			}
			l, h, cheating = 1, 10, false
		}
		if h < l {
			cheating = true
		}
	}
}
