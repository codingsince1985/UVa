// UVa 11854 - Egypt

package main

import (
	"fmt"
	"os"
)

func isRightTriangle(a, b, c int) bool {
	switch {
	case a >= b && a >= c:
		return b*b+c*c == a*a
	case b >= a && b >= c:
		return a*a+c*c == b*b
	default:
		return a*a+b*b == c*c
	}
}

func main() {
	in, _ := os.Open("11854.in")
	defer in.Close()
	out, _ := os.Create("11854.out")
	defer out.Close()

	var a, b, c int
	for {
		if fmt.Fscanf(in, "%d%d%d", &a, &b, &c); a == 0 && b == 0 && c == 0 {
			break
		}
		if isRightTriangle(a, b, c) {
			fmt.Fprintln(out, "right")
		} else {
			fmt.Fprintln(out, "wrong")
		}
	}
}
