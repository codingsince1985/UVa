// UVa 138 - Street Numbers

package main

import (
	"fmt"
	"os"
)

var n, x [11]int

func N(idx int) int {
	if n[idx] == 0 {
		switch idx {
		case 0:
			n[idx] = 1
		case 1:
			n[idx] = 6
		default:
			n[idx] = 6*N(idx-1) - N(idx-2)
		}
	}
	return n[idx]
}

func X(idx int) int {
	if x[idx] == 0 {
		if idx == 0 {
			x[idx] = 1
		} else {
			x[idx] = N(idx) + N(idx-1) + X(idx-1)
		}
	}
	return x[idx]
}

func main() {
	out, _ := os.Create("138.out")
	defer out.Close()

	for i := 1; i <= 10; i++ {
		fmt.Fprintf(out, "%10d%10d\n", N(i), X(i))
	}
}
