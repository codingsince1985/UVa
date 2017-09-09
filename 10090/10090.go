// UVa 10090 - Marbles

package main

import (
	"fmt"
	"os"
)

func extendedEuclidean(n1, n2 int) (int, int, int) {
	if n2 == 0 {
		return 1, 0, n1
	}
	x, y, d := extendedEuclidean(n2, n1%n2)
	x, y = y, x-n1/n2*y
	return x, y, d
}

func solve(out *os.File, n, c1, n1, c2, n2 int) {
	x, y, d := extendedEuclidean(n1, n2)
	if n%d != 0 {
		fmt.Fprintln(out, "failed")
		return
	}
	x *= n / d
	y *= n / d
	l := -x * d / n2
	if x*d%n2 != 0 && l >= 0 {
		l++
	}
	h := y * d / n1
	if y*d%n2 != 0 && h <= 0 {
		h--
	}
	if l > h {
		fmt.Fprintln(out, "failed")
		return
	}
	if c1*(x+n2/d*l)+c2*(y-n1/d*l) > c1*(x+n2/d*h)+c2*(y-n1/d*h) {
		l = h
	}
	fmt.Fprintln(out, x+n2/d*l, y-n1/d*l)
}

func main() {
	in, _ := os.Open("10090.in")
	defer in.Close()
	out, _ := os.Create("10090.out")
	defer out.Close()

	var n, c1, n1, c2, n2 int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fscanf(in, "%d%d", &c1, &n1)
		fmt.Fscanf(in, "%d%d", &c2, &n2)
		solve(out, n, c1, n1, c2, n2)
	}
}
